package helper

var validInvestigationStyles = []string{
	"Evidence-Based Investigation",
	"Interview-Based Investigation",
	"Undercover Investigation",
	"Follow The Money Investigation",
	"Report-Based Investigation",
}

func IsValidInvestigationStyle(style string) bool {
	for _, v := range validInvestigationStyles {
		if style == v {
			return true
		}
	}
	return false
}

var validCaseStatus = []string{
	"Open",
	"In Progress",
	"Closed",
}

func IsValidCaseStatus(status string) bool {
	for _, v := range validCaseStatus {
		if status == v {
			return true
		}
	}
	return false
}

var validSuspectStatus = []string{
	"Arrested",
	"Released",
	"Wanted",
	"Under Investigation",
	"Eyewitness",
}

func IsValidSuspectStatus(status string) bool {
	for _, v := range validSuspectStatus {
		if status == v {
			return true
		}
	}
	return false
}

var validSuspectGender = []string{
	"Male",
	"Female",
}

func IsValidSuspectGender(gender string) bool {
	for _, v := range validSuspectGender {
		if gender == v {
			return true
		}
	}
	return false
}
