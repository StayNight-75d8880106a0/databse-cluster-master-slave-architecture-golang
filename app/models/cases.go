package models

import (
	"time"

	"gorm.io/datatypes"
)

type Status string

const (
	OPEN        Status = "Open"
	IN_PROGRESS Status = "In Progress"
	CLOSED      Status = "Closed"
)

type Cases struct {
	ID               *string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Case_Number      *string `json:"case_number"`
	Case_Title       *string `json:"case_title"`
	Case_Description *string `json:"case_description"`
	Incident_Date    datatypes.Date
	Location         *string   `json:"location"`
	Status           Status    `json:"status" gorm:"type:varchar(255)"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`

	Suspects  []Suspects  `json:"suspects" gorm:"foreignKey:Case_ID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Detective []Detective `gorm:"many2many:case_detectives;joinForeignKey:CaseID;JoinReferences:DetectiveID;constraint:OnDelete:CASCADE;"`
}

func (Cases) TableName() string {
	return "cases"
}
