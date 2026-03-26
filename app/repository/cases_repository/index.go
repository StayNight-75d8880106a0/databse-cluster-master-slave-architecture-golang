package cases_repository

import (
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/models"

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

func (repo *Cases_Repository) GetAll(limit int, offset int) ([]models.Cases, int64, error) {

	var cases []models.Cases

	var count int64

	errCount := repo.slave.Table("cases").Count(&count).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	errGet := repo.slave.Table("cases").Preload("Suspects").Limit(limit).Offset(offset).Find(&cases).Error

	return cases, count, errGet

}

func (repo *Cases_Repository) GetById(ID string) (*models.Cases, error) {

	var cases *models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Where("id = ?", ID).First(&cases).Error

	return cases, errGet

}

func (repo *Cases_Repository) GetByCaseNumber(case_number string) (*models.Cases, error) {

	var cases *models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Where("case_number = ?", case_number).First(&cases).Error

	return cases, errGet

}

func (repo *Cases_Repository) Update(ID string, cases *models.Cases) error {

	errUpdate := repo.master.Table("cases").Where("id = ?", ID).Updates(cases).Error

	return errUpdate

}

func (repo *Cases_Repository) Delete(ID string) error {

	var cases *models.Cases

	errDelete := repo.master.Table("cases").Unscoped().Where("id = ?", ID).Delete(&cases).Error

	return errDelete

}
