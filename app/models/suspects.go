package models

import (
	"time"

	"gorm.io/datatypes"
)

type Gender string

type Status_Suspect string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

const (
	Arrested            Status_Suspect = "Arrested"
	Released            Status_Suspect = "Released"
	Wanted              Status_Suspect = "Wanted"
	Under_Investigation Status_Suspect = "Under Investigation"
	Eyewitness          Status_Suspect = "Eyewitness"
)

type Suspects struct {
	ID             *string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Case_ID        *string `json:"case_id" gorm:"type:uuid;not null"`
	ID_card_Number *string `json:"id_card_number"`
	Full_Name      *string `json:"full_name"`
	Gender         Gender  `json:"gender" gorm:"type:varchar(255)"`
	Date_Of_Birth  datatypes.Date
	Address        *string        `json:"address"`
	Phone          *string        `json:"phone"`
	Occupation     *string        `json:"occupation"`
	Alibi          *string        `json:"alibi"`
	Status         Status_Suspect `json:"status" gorm:"type:varchar(255)"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
}

func (Suspects) TableName() string {
	return "suspects"
}
