package cases_repository_interface

import "databse-cluster-master-slave-architecture-golang/app/models"

type Cases_Repository_Interface interface {
	Create(cases *models.Cases) error
	GetAll(limit int, offset int, search string) ([]models.Cases, int64, error)
	GetById(ID string) (*models.Cases, error)
	Update(ID string, cases *models.Cases) error
	Delete(ID string) error
	GetCount() (int64, int64, int64, int64, error)
	FindDetectiveByIdViaCases(id []string, detectives *[]models.Detective) error
	UpdateDetectiveRelation(ID string, detectives []models.Detective) error
	GetCasesLatestUpdate() ([]models.Cases, error)
}
