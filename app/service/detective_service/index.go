package detective_service

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/repository/detective_repository_interface"
	"databse-cluster-master-slave-architecture-golang/app/models"
	"databse-cluster-master-slave-architecture-golang/app/request/detective_request"
	"errors"
	"math"
	"time"

	"gorm.io/gorm"
)

type Detective_Service struct {
	repository detective_repository_interface.Detective_Repository_Interface
}

func NewDetectiveServiceRegistry(detective_repository detective_repository_interface.Detective_Repository_Interface) *Detective_Service {
	return &Detective_Service{
		repository: detective_repository,
	}
}

func (s *Detective_Service) Create(detective_dto *detective_request.Detective_Dto) (detective_request.Detective_Response, error) {

	if detective_dto.Name == nil || *detective_dto.Name == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Name Cannot Be Empty!")
	}

	if detective_dto.Badge_Number == nil || *detective_dto.Badge_Number == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Bagde Number Cannot Be Empty!")
	}

	if detective_dto.Department == nil || *detective_dto.Department == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Departement Cannot Be Empty!")
	}

	if detective_dto.Station == nil || *detective_dto.Station == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Station Cannot Be Empty!")
	}

	if detective_dto.Phone == nil || *detective_dto.Phone == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Phone Cannot Be Empty")
	}

	if detective_dto.Investigation_Style == nil || *detective_dto.Investigation_Style == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Investigation Style Cannot Be Empty!")
	}

	detective := &models.Detective{
		Name:                detective_dto.Name,
		Badge_Number:        detective_dto.Badge_Number,
		Department:          detective_dto.Department,
		Station:             detective_dto.Station,
		Phone:               detective_dto.Phone,
		Investigation_Style: models.InvestigationStyle(*detective_dto.Investigation_Style),
	}

	errCreate := s.repository.Create(detective)

	if errCreate != nil {
		return detective_request.Detective_Response{}, helper.NewInternalServerError("An error occurred while adding detective data :" + errCreate.Error())
	}

	response := &detective_request.Detective_Response{
		Name:                detective.Name,
		Badge_Number:        detective.Badge_Number,
		Department:          detective.Department,
		Station:             detective.Station,
		Phone:               detective.Phone,
		Investigation_Style: (*string)(&detective.Investigation_Style),
		CreatedAt:           detective.CreatedAt,
		UpdatedAt:           detective.UpdatedAt,
	}

	return *response, nil

}

func (s *Detective_Service) GetAll(page int, limit int) ([]detective_request.Detective_Response, helper.PaginationMeta, error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	detective, totalData, errGet := s.repository.GetAll(limit, offset)

	if errGet != nil {
		return []detective_request.Detective_Response{}, helper.PaginationMeta{}, helper.NewInternalServerError("An error occurred while get detective data :" + errGet.Error())
	}

	var responses []detective_request.Detective_Response

	for _, value := range detective {
		response := detective_request.Detective_Response{
			Name:                value.Name,
			Badge_Number:        value.Badge_Number,
			Department:          value.Department,
			Station:             value.Station,
			Phone:               value.Phone,
			Investigation_Style: (*string)(&value.Investigation_Style),
			CreatedAt:           value.CreatedAt,
			UpdatedAt:           value.UpdatedAt,
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

func (s *Detective_Service) GetById(ID string) (detective_request.Detective_Response, error) {

	detective, errGet := s.repository.GetByID(ID)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return detective_request.Detective_Response{}, helper.NewNotFound("An error occurred while get detective data : " + errGet.Error())
		}
	}

	reponse := detective_request.Detective_Response{
		Name:                detective.Name,
		Badge_Number:        detective.Badge_Number,
		Department:          detective.Department,
		Station:             detective.Station,
		Phone:               detective.Phone,
		Investigation_Style: (*string)(&detective.Investigation_Style),
		CreatedAt:           detective.CreatedAt,
		UpdatedAt:           detective.UpdatedAt,
	}

	return reponse, nil
}

func (s *Detective_Service) Update(ID string, detective_dto *detective_request.Detective_Dto) (detective_request.Detective_Response, error) {

	GetDetective, errGet := s.repository.GetByID(ID)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return detective_request.Detective_Response{}, helper.NewNotFound("An error occurred while get detective data : " + errGet.Error())
		}
	}

	if detective_dto.Name == nil || *detective_dto.Name == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Name Cannot Be Empty!")
	}

	if detective_dto.Badge_Number == nil || *detective_dto.Badge_Number == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Bagde Number Cannot Be Empty!")
	}

	if detective_dto.Department == nil || *detective_dto.Department == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Departement Cannot Be Empty!")
	}

	if detective_dto.Station == nil || *detective_dto.Station == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Station Cannot Be Empty!")
	}

	if detective_dto.Phone == nil || *detective_dto.Phone == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Phone Cannot Be Empty")
	}

	if detective_dto.Investigation_Style == nil || *detective_dto.Investigation_Style == "" {
		return detective_request.Detective_Response{}, helper.NewBadRequest("Investigation Style Cannot Be Empty!")
	}

	detective := &models.Detective{
		Name:                detective_dto.Name,
		Badge_Number:        detective_dto.Badge_Number,
		Department:          detective_dto.Department,
		Station:             detective_dto.Station,
		Phone:               detective_dto.Phone,
		Investigation_Style: models.InvestigationStyle(*detective_dto.Investigation_Style),
		CreatedAt:           GetDetective.CreatedAt,
		UpdatedAt:           time.Now(),
	}

	errUpdate := s.repository.Update(ID, detective)

	if errUpdate != nil {
		return detective_request.Detective_Response{}, helper.NewInternalServerError("An error occurred while update detective data :" + errUpdate.Error())
	}

	response := &detective_request.Detective_Response{
		Name:                detective.Name,
		Badge_Number:        detective.Badge_Number,
		Department:          detective.Badge_Number,
		Station:             detective.Station,
		Phone:               detective.Phone,
		Investigation_Style: (*string)(&detective.Investigation_Style),
		CreatedAt:           detective.CreatedAt,
		UpdatedAt:           detective.UpdatedAt,
	}

	return *response, nil
}

func (s *Detective_Service) Delete(ID string) error {

	_, errGet := s.repository.GetByID(ID)

	if errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			return helper.NewNotFound("An error occurred while get detective data : " + errGet.Error())
		}
	}

	errDelete := s.repository.Delete(ID)

	if errDelete != nil {
		helper.NewInternalServerError("An error occurred while delete detective data :" + errDelete.Error())
	}

	return nil

}
