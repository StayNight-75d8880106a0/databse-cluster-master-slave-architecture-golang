package detective_service_interface

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/request/detective_request"
)

type Detective_Service_Interface interface {
	Create(detective_dto *detective_request.Detective_Dto) (detective_request.Detective_Response, error)
	GetAll(page int, limit int) ([]detective_request.Detective_Response, helper.PaginationMeta, error)
	GetById(ID string) (detective_request.Detective_Response, error)
	Update(ID string, detective_dto *detective_request.Detective_Dto) (detective_request.Detective_Response, error)
	Delete(ID string) error
}
