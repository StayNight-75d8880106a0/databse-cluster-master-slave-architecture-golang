package cases_repository_interface

import "databse-cluster-master-slave-architecture-golang/app/models"

type Cases_Repository_Interface interface {
	Create(cases *models.Cases) error
	GetAll(limit int, offset int) ([]models.Cases, int64, error)
	GetById(ID string) (*models.Cases, error)
	Update(ID string, cases *models.Cases) error
	Delete(ID string) error
	GetByCaseNumber(case_number string) (*models.Cases, error)
}
