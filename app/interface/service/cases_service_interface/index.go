package cases_service_interface

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/request/cases_request"
)

type Cases_Service_Interface interface {
	Create(cases_dto *cases_request.Cases_Dto) (cases_request.Cases_Response, error)
	GetAll(page int, limit int, search string) ([]cases_request.Cases_Response, helper.PaginationMeta, error)
	GetById(ID string) (cases_request.Cases_Response, error)
	Update(ID string, cases_dto *cases_request.Cases_Dto) (cases_request.Cases_Response, error)
	Delete(ID string) error
	GetCount() (*cases_request.Cases_Count, error)
	GetCasesLatestUpdate() ([]cases_request.Cases_Response, error)
}
