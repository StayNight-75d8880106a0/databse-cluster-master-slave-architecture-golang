package detective_repository

import (
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/models"

	"gorm.io/gorm"
)

type Detective_Repository struct {
	Master *gorm.DB
	Slave  *gorm.DB
}

func NewDetectiveRepositoryRegistry() *Detective_Repository {
	dbCluster := database.GetInstanceDbCluster()
	return &Detective_Repository{
		Master: dbCluster.Master,
		Slave:  dbCluster.SlaveDetective,
	}
}

func (repo *Detective_Repository) Create(detective *models.Detective) error {

	errCreate := repo.Master.Table("detective").Create(&detective).Error

	return errCreate

}

func (repo *Detective_Repository) GetAll(limit int, offset int) ([]models.Detective, int64, error) {

	var detective []models.Detective

	var count int64

	errCount := repo.Slave.Table("detective").Count(&count).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	errGet := repo.Slave.Table("detective").Preload("Cases").Limit(limit).Offset(offset).Find(&detective).Error

	return detective, count, errGet

}

func (repo *Detective_Repository) GetByID(ID string) (*models.Detective, error) {

	var detective *models.Detective

	errGet := repo.Slave.Table("detective").Preload("Cases").Where("id = ?", ID).First(&detective).Error

	return detective, errGet

}

func (repo *Detective_Repository) Update(ID string, detective *models.Detective) error {

	errUpdate := repo.Master.Table("detective").Where("id = ?", ID).Updates(detective).Error

	return errUpdate

}

func (repo *Detective_Repository) Delete(ID string) error {

	var detective *models.Detective

	errDelete := repo.Master.Table("detective").Unscoped().Where("id = ?", ID).Delete(&detective).Error

	return errDelete

}
