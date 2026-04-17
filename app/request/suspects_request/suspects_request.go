package suspects_request

import "time"

type Suspects_Request struct {
	ID_Card_Number string `json:"id_card_number"`
	Full_Name      string `json:"full_name"`
	Gender         string `json:"gender" binding:"required,oneof='Male' 'Female'"`
	Date_Of_Birth  string `json:"date_of_birth"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	Occupation     string `json:"occupation"`
	Alibi          string `json:"alibi"`
	Status         string `json:"status" binding:"required,oneof='Arrested' 'Released' 'Wanted' 'Under Investigation' 'Eyewitness'"`
}

type Suspects_Dto struct {
	Case_ID        *string
	ID_Card_Number *string
	Full_Name      *string
	Gender         *string
	Date_Of_Birth  time.Time
	Address        *string
	Phone          *string
	Occupation     *string
	Alibi          *string
	Status         *string
}

type Suspects_Response struct {
	Case_ID        *string   `json:"case_id"`
	ID_card_Number *string   `json:"id_card_number"`
	Full_Name      *string   `json:"full_name"`
	Gender         *string   `json:"gender"`
	Date_Of_Birth  time.Time `json:"date_of_birth"`
	Address        *string   `json:"address"`
	Phone          *string   `json:"phone"`
	Occupation     *string   `json:"occupation"`
	Alibi          *string   `json:"alibi"`
	Status         *string   `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
