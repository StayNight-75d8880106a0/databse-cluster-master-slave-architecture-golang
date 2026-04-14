package detective_repository_interface

import "databse-cluster-master-slave-architecture-golang/app/models"

type Detective_Repository_Interface interface {
	Create(detective *models.Detective) error
	GetAll(limit int, offset int) ([]models.Detective, int64, error)
	GetByID(ID string) (*models.Detective, error)
	Update(ID string, detective *models.Detective) error
	Delete(ID string) error
}
