package models

type CaseDetective struct {
	CaseID      string `gorm:"column:case_id;type:uuid"`
	DetectiveID string `gorm:"column:detective_id;type:uuid"`
}
