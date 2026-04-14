package detective_request

import "time"

type Detective_Request struct {
	Name                *string `json:"name"`
	Badge_Number        *string `json:"badge_number"`
	Department          *string `json:"department"`
	Station             *string `json:"station"`
	Phone               *string `json:"phone"`
	Investigation_Style *string `json:"investigation_style" binding:"required,oneof='Evidence-Based Investigation' 'Interview-Based Investigation' 'Undercover Investigation' 'Follow The Money Investigation' 'Report-Based Investigation'"`
}

type Detective_Dto struct {
	Name                *string
	Badge_Number        *string
	Department          *string
	Station             *string
	Phone               *string
	Investigation_Style *string
}

type Detective_Response struct {
	Name                *string   `json:"name"`
	Badge_Number        *string   `json:"badge_number"`
	Department          *string   `json:"department"`
	Station             *string   `json:"station"`
	Phone               *string   `json:"phone"`
	Investigation_Style *string   `json:"investigation_style" binding:"required,oneof='Evidence-Based Investigation' 'Interview-Based Investigation' 'Undercover Investigation' 'Follow The Money Investigation' 'Report-Based Investigation'"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
