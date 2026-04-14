package models

import "time"

type InvestigationStyle string

const (
	Evidence_Based_Investigation   InvestigationStyle = "Evidence-Based Investigation"
	Interview_Based_Investigation  InvestigationStyle = "Interview-Based Investigation"
	Undercover_Investigation       InvestigationStyle = "Undercover Investigation"
	Follow_The_Money_Investigation InvestigationStyle = "Follow The Money Investigation"
	Report_Based_Investigation     InvestigationStyle = "Report-Based Investigation"
)

type Detective struct {
	ID                  *string            `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name                *string            `json:"name"`
	Badge_Number        *string            `json:"badge_number"`
	Department          *string            `json:"department"`
	Station             *string            `json:"station"`
	Phone               *string            `json:"phone"`
	Investigation_Style InvestigationStyle `json:"investigation_style" gorm:"type:varchar(255)"`
	CreatedAt           time.Time          `gorm:"column:created_at"`
	UpdatedAt           time.Time          `gorm:"column:updated_at"`

	Cases []Cases `gorm:"many2many:case_detectives;joinForeignKey:DetectiveID;JoinReferences:CaseID;"`
}

func (Detective) TableName() string {
	return "detective"
}
