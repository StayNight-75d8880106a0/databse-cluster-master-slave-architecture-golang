package cases_repository

import (
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/models"
	"fmt"

	"gorm.io/gorm"
)

type Cases_Repository struct {
	master *gorm.DB
	slave  *gorm.DB
}

func NewCasesRepositoryRegistry() *Cases_Repository {
	dbCluster := database.GetInstanceDbCluster()
	return &Cases_Repository{
		master: dbCluster.Master,
		slave:  dbCluster.SlaveCases,
	}
}

func (repo *Cases_Repository) Create(cases *models.Cases) error {

	errCreate := repo.master.Table("cases").Create(cases).Error

	return errCreate

}

func (repo *Cases_Repository) GetAll(limit int, offset int, search string) ([]models.Cases, int64, error) {

	var cases []models.Cases

	var count int64

	baseQuery := repo.slave.Table("cases")

	if search != "" {
		baseQuery = baseQuery.Where("case_title LIKE ?", search+"%")
	}

	queryCount := baseQuery.Session(&gorm.Session{})

	errCount := queryCount.Count(&count).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	queryGet := baseQuery.Session(&gorm.Session{})

	errGet := queryGet.Preload("Suspects").Preload("Detective").Limit(limit).Offset(offset).Find(&cases).Error

	return cases, count, errGet

}

func (repo *Cases_Repository) GetCount() (int64, int64, int64, int64, error) {

	var TotalCases int64
	var TotalOpenCases int64
	var TotalClosedCases int64
	var TotalPendingCases int64

	errCountCases := repo.slave.Table("cases").Count(&TotalCases).Error
	errCountOpenCases := repo.slave.Table("cases").Where("status = ?", "Open").Count(&TotalOpenCases).Error
	errCountClosedCases := repo.slave.Table("cases").Where("status = ?", "Closed").Count(&TotalClosedCases).Error
	errCountPendingCases := repo.slave.Table("cases").Where("status = ?", "In Progress").Count(&TotalPendingCases).Error

	if errCountCases != nil || errCountOpenCases != nil || errCountClosedCases != nil || errCountPendingCases != nil {
		return 0, 0, 0, 0, fmt.Errorf("error occurred while counting cases")
	}

	return TotalCases, TotalOpenCases, TotalClosedCases, TotalPendingCases, nil

}

func (repo *Cases_Repository) GetById(ID string) (*models.Cases, error) {

	var cases *models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Preload("Detective").Where("id = ?", ID).First(&cases).Error

	return cases, errGet

}

func (repo *Cases_Repository) Update(ID string, cases *models.Cases) error {

	errUpdate := repo.master.Table("cases").Where("id = ?", ID).Updates(cases).Error

	return errUpdate

}

func (repo *Cases_Repository) Delete(ID string) error {

	return repo.master.Transaction(func(tx *gorm.DB) error {
		// Remove many-to-many rows first to satisfy FK constraints.
		errDeleteRelation := tx.Table("case_detectives").Where("case_id = ?", ID).Delete(&models.CaseDetective{}).Error
		if errDeleteRelation != nil {
			return errDeleteRelation
		}

		var cases models.Cases
		errDelete := tx.Table("cases").Unscoped().Where("id = ?", ID).Delete(&cases).Error
		if errDelete != nil {
			return errDelete
		}

		return nil
	})

}

func (repo *Cases_Repository) FindDetectiveByIdViaCases(id []string, detectives *[]models.Detective) error {

	errGet := repo.master.Table("detective").Where("id IN ?", id).Find(detectives).Error

	return errGet
}

func (repo *Cases_Repository) UpdateDetectiveRelation(ID string, detectives []models.Detective) error {

	var cases models.Cases

	err := repo.master.First(&cases, "id = ?", ID).Error
	if err != nil {
		return err
	}

	errUpdate := repo.master.Model(&cases).Association("Detective").Replace(detectives)

	return errUpdate

}

func (repo *Cases_Repository) GetCasesLatestUpdate() ([]models.Cases, error) {

	var cases []models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Preload("Detective").Order("updated_at DESC").Limit(5).Find(&cases).Error

	return cases, errGet

}
