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

	if len(cases_dto.DetectiveIDs) == 0 {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Detective Cannot Be Empty!")
	}

	if !helper.IsValidCaseStatus(*cases_dto.Status) {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Invalid Case Status")
	}

	year := time.Now().Year()
	month := time.Now().Month()

	caseNumber := fmt.Sprintf("CASE-%s-%d-%s", RandomNumber, year, month)

	var detectives []models.Detective

	errGetDetective := s.repository.FindDetectiveByIdViaCases(cases_dto.DetectiveIDs, &detectives)

	if errGetDetective != nil {
		return cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while adding detective id data : " + errGetDetective.Error())
	}

	cases := &models.Cases{
		Case_Number:      &caseNumber,
		Case_Title:       cases_dto.Case_Title,
		Case_Description: cases_dto.Case_Description,
		Incident_Date:    datatypes.Date(cases_dto.Incident_Date),
		Location:         cases_dto.Location,
		Status:           models.Status(*cases_dto.Status),
		Detective:        detectives,
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
		Status:           (*string)(&cases.Status),
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         cases.Suspects,
	}

	if errCreate == nil {
		helper.Manager.Broadcast(map[string]interface{}{
			"event": "CASES_CREATED",
			"data":  response,
		})
	}

	return *response, nil

}

func (s *Cases_Service) GetAll(page int, limit int, search string) ([]cases_request.Cases_Response, helper.PaginationMeta, error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	cases, totalData, errGet := s.repository.GetAll(limit, offset, search)

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
			Status:           (*string)(&value.Status),
			CreatedAt:        value.CreatedAt,
			UpdatedAt:        value.UpdatedAt,
			Suspects:         value.Suspects,
			Detective:        value.Detective,
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

func (s *Cases_Service) GetCount() (*cases_request.Cases_Count, error) {

	TotalCases, TotalOpenCases, TotalClosedCases, TotalPendingCases, errGet := s.repository.GetCount()

	if errGet != nil {
		return &cases_request.Cases_Count{}, helper.NewInternalServerError("An error occurred while count case data :" + errGet.Error())
	}

	responses := &cases_request.Cases_Count{
		TotalCases:        TotalCases,
		TotalOpenCases:    TotalOpenCases,
		TotalClosedCases:  TotalClosedCases,
		TotalPendingCases: TotalPendingCases,
	}

	return responses, nil

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
		Status:           (*string)(&cases.Status),
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         cases.Suspects,
		Detective:        cases.Detective,
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

	if len(cases_dto.DetectiveIDs) == 0 {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Detective Cannot Be Empty!")
	}

	var detectives []models.Detective

	errGetDetective := s.repository.FindDetectiveByIdViaCases(cases_dto.DetectiveIDs, &detectives)

	if errGetDetective != nil {
		return cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while adding detective id data : " + errGetDetective.Error())
	}

	if !helper.IsValidCaseStatus(*cases_dto.Status) {
		return cases_request.Cases_Response{}, helper.NewBadRequest("Invalid Case Status")
	}

	cases := &models.Cases{
		Case_Title:       cases_dto.Case_Title,
		Case_Description: cases_dto.Case_Description,
		Incident_Date:    datatypes.Date(cases_dto.Incident_Date),
		Location:         cases_dto.Location,
		Status:           models.Status(*cases_dto.Status),
	}

	errUpdate := s.repository.Update(ID, cases)

	if errUpdate != nil {
		return cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while update case data : " + errUpdate.Error())
	}

	errRelation := s.repository.UpdateDetectiveRelation(ID, detectives)

	if errRelation != nil {
		return cases_request.Cases_Response{}, helper.NewInternalServerError("Failed to update detectives relation: " + errRelation.Error())
	}

	response := &cases_request.Cases_Response{
		Case_Number:      GetCases.Case_Number,
		Case_Title:       cases.Case_Title,
		Case_Description: cases.Case_Description,
		Incident_Date:    time.Time(cases.Incident_Date),
		Location:         cases.Location,
		Status:           (*string)(&cases.Status),
		CreatedAt:        cases.CreatedAt,
		UpdatedAt:        cases.UpdatedAt,
		Suspects:         GetCases.Suspects,
	}

	if errUpdate == nil {
		helper.Manager.Broadcast(map[string]interface{}{
			"event": "CASES_UPDATED",
			"data":  response,
		})
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

	if errDelete == nil {
		helper.Manager.Broadcast(map[string]interface{}{
			"event":   "CASE_DELETED",
			"message": "Success Delete Case",
		})
	}

	return nil

}

func (s *Cases_Service) GetCasesLatestUpdate() ([]cases_request.Cases_Response, error) {

	cases, errGet := s.repository.GetCasesLatestUpdate()

	if errGet != nil {
		return []cases_request.Cases_Response{}, helper.NewInternalServerError("An error occurred while get case data : " + errGet.Error())
	}

	var responses []cases_request.Cases_Response

	for _, value := range cases {
		response := cases_request.Cases_Response{
			Case_Number:      value.Case_Number,
			Case_Title:       value.Case_Title,
			Case_Description: value.Case_Description,
			Incident_Date:    time.Time(value.Incident_Date),
			Location:         value.Location,
			Status:           (*string)(&value.Status),
			CreatedAt:        value.CreatedAt,
			UpdatedAt:        value.UpdatedAt,
			Suspects:         value.Suspects,
			Detective:        value.Detective,
		}
		responses = append(responses, response)
	}

	return responses, nil

}
