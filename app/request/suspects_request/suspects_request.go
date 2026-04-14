package suspects_request

import "time"

type Suspects_Request struct {
	ID_Card_Number *string `json:"id_card_number"`
	Full_Name      *string `json:"full_name"`
	Address        *string `json:"address"`
	Alibi          *string `json:"alibi"`
}

type Suspects_Dto struct {
	Case_ID        *string
	ID_Card_Number *string
	Full_Name      *string
	Address        *string
	Alibi          *string
}

type Suspects_Response struct {
	Case_ID        *string   `json:"case_id"`
	ID_card_Number *string   `json:"id_card_number"`
	Full_Name      *string   `json:"full_name"`
	Address        *string   `json:"address"`
	Alibi          *string   `json:"alibi"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
