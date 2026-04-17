package cases_request

import (
	"databse-cluster-master-slave-architecture-golang/app/models"
	"time"
)

type Cases_Request struct {
	Case_Title       string   `json:"case_title"`
	Case_Description string   `json:"case_description"`
	Incident_Date    string   `json:"incident_date"`
	Location         string   `json:"location"`
	Status           string   `json:"status" binding:"required,oneof='Open' 'In Progress' 'Closed'"`
	DetectiveIDs     []string `json:"detective_ids"`
}

type Cases_Dto struct {
	Case_Title       *string
	Case_Description *string
	Incident_Date    time.Time
	Location         *string
	Status           *string
	DetectiveIDs     []string
}

type Cases_Response struct {
	Case_Number      *string            `json:"case_number"`
	Case_Title       *string            `json:"case_title"`
	Case_Description *string            `json:"case_description"`
	Incident_Date    time.Time          `json:"incident_date"`
	Location         *string            `json:"location"`
	Status           *string            `json:"status" binding:"required,oneof='Open' 'In Progress' 'Closed'"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	Suspects         []models.Suspects  `json:"suspects"`
	Detective        []models.Detective `json:"detectives"`
}

type Cases_Count struct {
	TotalCases        int64 `json:"total_cases"`
	TotalOpenCases    int64 `json:"total_open_cases"`
	TotalClosedCases  int64 `json:"total_closed_cases"`
	TotalPendingCases int64 `json:"total_pending_cases"`
}
