package suspect_service

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/repository/suspect_repository_interface"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/cases_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/models"
	"databse-cluster-master-slave-architecture-golang/app/request/suspects_request"
	"errors"
	"math"

	"gorm.io/gorm"
)

type Suspect_Service struct {
	repository  suspect_repository_interface.Suspect_Repository_Interface
	CaseService cases_service_interface.Cases_Service_Interface
}

func NewSuspectServiceRegistry(suspect_repository suspect_repository_interface.Suspect_Repository_Interface,
	case_service cases_service_interface.Cases_Service_Interface) *Suspect_Service {
	return &Suspect_Service{
		repository:  suspect_repository,
		CaseService: case_service,
	}
}

func (s *Suspect_Service) Create(ID_Case string, suspect_dto *suspects_request.Suspects_Dto) (suspects_request.Suspects_Response, error) {

	if suspect_dto.Case_ID == nil || *suspect_dto.Case_ID == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Case Id Cannot Be Empty")
	}

	_, errGet := s.CaseService.GetById(ID_Case)

	if errGet != nil {
		return suspects_request.Suspects_Response{}, errGet
	}

	if suspect_dto.ID_Card_Number == nil || *suspect_dto.ID_Card_Number == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Id Card Number Cannot Be Empty")
	}

	if suspect_dto.Full_Name == nil || *suspect_dto.Full_Name == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Full Name Cannot Be Empty")
	}

	if suspect_dto.Address == nil || *suspect_dto.Address == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Address Cannot Be Empty")
	}

	if suspect_dto.Alibi == nil || *suspect_dto.Alibi == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Alibi Cannot Be Empty")
	}

	suspect := &models.Suspects{
		Case_ID:        suspect_dto.Case_ID,
		ID_card_Number: suspect_dto.ID_Card_Number,
		Full_Name:      suspect_dto.Full_Name,
		Address:        suspect_dto.Address,
		Alibi:          suspect_dto.Alibi,
	}

	errCreate := s.repository.Create(suspect)

	if errCreate != nil {
		return suspects_request.Suspects_Response{}, helper.NewInternalServerError("An error occurred while adding suspect data : " + errCreate.Error())
	}

	response := &suspects_request.Suspects_Response{
		Case_ID:        suspect.Case_ID,
		ID_card_Number: suspect.ID_card_Number,
		Full_Name:      suspect.Full_Name,
		Address:        suspect.Address,
		Alibi:          suspect.Alibi,
		CreatedAt:      suspect.CreatedAt,
		UpdatedAt:      suspect.UpdatedAt,
	}

	return *response, nil

}

func (s *Suspect_Service) GetAll(ID_Case string, page int, limit int) ([]suspects_request.Suspects_Response, helper.PaginationMeta, error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	suspect, totalData, errGet := s.repository.GetAll(ID_Case, limit, offset)

	if errGet != nil {
		return []suspects_request.Suspects_Response{}, helper.PaginationMeta{}, helper.NewInternalServerError("An error occurred while get suspect data : " + errGet.Error())
	}

	var responses []suspects_request.Suspects_Response

	for _, value := range suspect {
		response := &suspects_request.Suspects_Response{
			Case_ID:        value.Case_ID,
			ID_card_Number: value.ID_card_Number,
			Full_Name:      value.Full_Name,
			Address:        value.Address,
			Alibi:          value.Alibi,
			CreatedAt:      value.CreatedAt,
			UpdatedAt:      value.UpdatedAt,
		}
		responses = append(responses, *response)
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

func (s *Suspect_Service) GetById(ID string, ID_Case string) (suspects_request.Suspects_Response, error) {

	suspect, errGet := s.repository.GetById(ID, ID_Case)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return suspects_request.Suspects_Response{}, helper.NewNotFound("An error occurred while get case data : " + errGet.Error())
		}
	}

	response := &suspects_request.Suspects_Response{
		Case_ID:        suspect.Case_ID,
		ID_card_Number: suspect.ID_card_Number,
		Full_Name:      suspect.Full_Name,
		Address:        suspect.Address,
		Alibi:          suspect.Alibi,
	}

	return *response, nil

}

func (s *Suspect_Service) Update(ID string, ID_Case string, suspect_dto *suspects_request.Suspects_Dto) (suspects_request.Suspects_Response, error) {

	Getsuspect, errGet := s.repository.GetById(ID, ID_Case)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return suspects_request.Suspects_Response{}, helper.NewNotFound("An error occurred while get case data : " + errGet.Error())
		}
	}

	if suspect_dto.Case_ID == nil || *suspect_dto.Case_ID == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Case Id Cannot Be Empty")
	}

	_, errGetCase := s.CaseService.GetById(ID_Case)

	if errGetCase != nil {
		return suspects_request.Suspects_Response{}, errGetCase
	}

	if suspect_dto.ID_Card_Number == nil || *suspect_dto.ID_Card_Number == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Id Card Number Cannot Be Empty")
	}

	if suspect_dto.Full_Name == nil || *suspect_dto.Full_Name == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Full Name Cannot Be Empty")
	}

	if suspect_dto.Address == nil || *suspect_dto.Address == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Address Cannot Be Empty")
	}

	if suspect_dto.Alibi == nil || *suspect_dto.Alibi == "" {
		return suspects_request.Suspects_Response{}, helper.NewBadRequest("Alibi Cannot Be Empty")
	}

	suspect := &models.Suspects{
		ID_card_Number: suspect_dto.ID_Card_Number,
		Full_Name:      suspect_dto.Full_Name,
		Address:        suspect_dto.Address,
		Alibi:          suspect_dto.Alibi,
	}

	errUpdate := s.repository.Update(ID, ID_Case, suspect)

	if errUpdate != nil {
		return suspects_request.Suspects_Response{}, helper.NewInternalServerError("An error occurred while update suspect data : " + errUpdate.Error())
	}

	response := &suspects_request.Suspects_Response{
		Case_ID:        Getsuspect.Case_ID,
		ID_card_Number: suspect.ID_card_Number,
		Full_Name:      suspect.Full_Name,
		Address:        suspect.Address,
		Alibi:          suspect.Alibi,
		CreatedAt:      Getsuspect.CreatedAt,
		UpdatedAt:      suspect.UpdatedAt,
	}

	return *response, nil

}

func (s *Suspect_Service) Delete(ID string, ID_Case string) error {

	_, errGet := s.repository.GetById(ID, ID_Case)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return helper.NewNotFound("An error occurred while get case data : " + errGet.Error())
		}
	}

	_, errGetCase := s.CaseService.GetById(ID_Case)

	if errGetCase != nil {
		return errGetCase
	}

	errDelete := s.repository.Delete(ID, ID_Case)

	if errDelete != nil {
		return helper.NewInternalServerError("An error occurred while update suspect data : " + errDelete.Error())
	}

	return nil

}
