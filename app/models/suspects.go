package models

import "time"

type Suspects struct {
	ID             *string   `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Case_ID        *string   `json:"case_id" gorm:"type:uuid;not null"`
	ID_card_Number *string   `json:"id_card_number"`
	Full_Name      *string   `json:"full_name"`
	Address        *string   `json:"address"`
	Alibi          *string   `json:"alibi"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (Suspects) TableName() string {
	return "suspects"
}
