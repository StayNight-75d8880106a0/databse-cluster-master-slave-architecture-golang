package suspect_service_interface

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/request/suspects_request"
)

type Suspect_Service_Interface interface {
	Create(ID_Case string, suspect_dto *suspects_request.Suspects_Dto) (suspects_request.Suspects_Response, error)
	GetAll(ID_Case string, page int, limit int) ([]suspects_request.Suspects_Response, helper.PaginationMeta, error)
	GetById(ID string, ID_Case string) (suspects_request.Suspects_Response, error)
	Update(ID string, ID_Case string, suspect_dto *suspects_request.Suspects_Dto) (suspects_request.Suspects_Response, error)
	Delete(ID string, ID_Case string) error
}
