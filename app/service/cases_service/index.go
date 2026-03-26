package cases_service

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/repository/cases_repository_interface"
	"databse-cluster-master-slave-architecture-golang/app/models"
	"databse-cluster-master-slave-architecture-golang/app/request/cases_request"
	"errors"
	"fmt"
	"math"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Cases_Service struct {
	repository cases_repository_interface.Cases_Repository_Interface
}

func NewCasesServiceRegistry(cases_repository cases_repository_interface.Cases_Repository_Interface) *Cases_Service {
	return &Cases_Service{
		repository: cases_repository,
	}
}

func (s *Cases_Service) Create(cases_dto *cases_request.Cases_Dto) (cases_request.Cases_Response, error) {

	RandomNumber := helper.GenerateRandomNumber()

	if cases_dto.Case_Title == nil || *cases_dto.Case_Title == "" {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Case Title Cannot Be Empty!")
	}

	if cases_dto.Case_Description == nil || *cases_dto.Case_Description == "" {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Case Description Cannot Be Empty!")
	}

	if cases_dto.Incident_Date.Local().IsZero() {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Incident Date Cannot Be Empty!")
	}

	if cases_dto.Location == nil || *cases_dto.Location == "" {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Location Cannot Be Empty!")
	}

	year := time.Now().Year()
	month := time.Now().Month()

	caseNumber := fmt.Sprintf("CASE-%s-%d-%s", RandomNumber, year, month)

	cases := &models.Cases{
		Case_Number:      &caseNumber,
		Case_Title:       cases_dto.Case_Title,
		Case_Description: cases_dto.Case_Description,
		Incident_Date:    datatypes.Date(cases_dto.Incident_Date),
		Location:         cases_dto.Location,
	}

	errCreate := s.repository.Create(cases)

	if errCreate != nil {
		return cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while adding case data : " + errCreate.Error())
	}

	response := &cases_request.Cases_Response{
		Case_Number:      cases.Case_Number,
		Case_Title:       cases.Case_Title,
		Case_Description: cases.Case_Description,
		Incident_Date:    time.Time(cases.Incident_Date),
		Location:         cases.Location,
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         cases.Suspects,
	}

	return *response, nil

}

func (s *Cases_Service) GetAll(page int, limit int) ([]cases_request.Cases_Response, helper.PaginationMeta, error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	cases, totalData, errGet := s.repository.GetAll(limit, offset)

	if errGet != nil {
		return []cases_request.Cases_Response{}, helper.PaginationMeta{}, helper.NewInternalServerError("An error occurred while get case data : " + errGet.Error())
	}

	var responses []cases_request.Cases_Response

	for _, value := range cases {
		response := cases_request.Cases_Response{
			Case_Number:      value.Case_Number,
			Case_Title:       value.Case_Title,
			Case_Description: value.Case_Description,
			Incident_Date:    time.Time(value.Incident_Date),
			Location:         value.Location,
			CreatedAt:        value.CreatedAt,
			UpdatedAt:        value.UpdatedAt,
			Suspects:         value.Suspects,
		}
		responses = append(responses, response)
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	meta := helper.PaginationMeta{
		Page:      page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return responses, meta, nil
}

func (s *Cases_Service) GetById(ID string) (cases_request.Cases_Response, error) {

	cases, errGet := s.repository.GetById(ID)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return cases_request.Cases_Response{}, helper.NewNotFound("An error occurred while get case data : " + errGet.Error())
		}
	}

	responses := &cases_request.Cases_Response{
		Case_Number:      cases.Case_Number,
		Case_Title:       cases.Case_Title,
		Case_Description: cases.Case_Description,
		Incident_Date:    time.Time(cases.Incident_Date),
		Location:         cases.Location,
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         cases.Suspects,
	}

	return *responses, nil

}

func (s *Cases_Service) GetByCaseNumber(case_number string) (cases_request.Cases_Response, error) {

	cases, errGet := s.repository.GetByCaseNumber(case_number)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return cases_request.Cases_Response{}, helper.NewNotFound("An error occurred while get case data : " + errGet.Error())
		}
	}

	responses := &cases_request.Cases_Response{
		Case_Number:      cases.Case_Number,
		Case_Title:       cases.Case_Title,
		Case_Description: cases.Case_Description,
		Incident_Date:    time.Time(cases.Incident_Date),
		Location:         cases.Location,
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         cases.Suspects,
	}

	return *responses, nil

}

func (s *Cases_Service) Update(ID string, cases_dto *cases_request.Cases_Dto) (cases_request.Cases_Response, error) {

	GetCases, errGet := s.repository.GetById(ID)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while get case data : " + errGet.Error())
		}
	}

	if cases_dto.Case_Title == nil || *cases_dto.Case_Title == "" {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Case Title Cannot Be Empty!")
	}

	if cases_dto.Case_Description == nil || *cases_dto.Case_Description == "" {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Case Description Cannot Be Empty!")
	}

	if cases_dto.Incident_Date.Local().IsZero() {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Incident Date Cannot Be Empty!")
	}

	if cases_dto.Location == nil || *cases_dto.Location == "" {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Location Cannot Be Empty!")
	}

	cases := &models.Cases{
		Case_Title:       cases_dto.Case_Title,
		Case_Description: cases_dto.Case_Description,
		Incident_Date:    datatypes.Date(cases_dto.Incident_Date),
		Location:         cases_dto.Location,
	}

	errUpdate := s.repository.Update(ID, cases)

	if errUpdate != nil {
		return cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while update case data : " + errUpdate.Error())
	}

	response := &cases_request.Cases_Response{
		Case_Number:      GetCases.Case_Number,
		Case_Title:       cases.Case_Title,
		Case_Description: cases.Case_Description,
		Incident_Date:    time.Time(cases.Incident_Date),
		Location:         cases.Location,
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         GetCases.Suspects,
	}

	return *response, nil

}

func (s *Cases_Service) Delete(ID string) error {

	_, errGet := s.repository.GetById(ID)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return helper.NewInternalServerError("An error occurred while get case data : " + errGet.Error())
		}
	}

	errDelete := s.repository.Delete(ID)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while delete case data : " + errDelete.Error())
	}

	return nil

}
