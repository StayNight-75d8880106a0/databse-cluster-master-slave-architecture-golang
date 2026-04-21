package vector_store

import (
	"context"
	"fmt"
	"time"

	"databse-cluster-master-slave-architecture-golang/app/models"

	"github.com/tmc/langchaingo/schema"
	"gorm.io/gorm"
)

func (v *AI_VectorStore) LoadDatabaseSnapshot(db *gorm.DB) error {
	ctx := context.Background()
	var docs []schema.Document

	// ==================== CASES ====================
	var cases []models.Cases
	db.Preload("Detective").Preload("Suspects").Find(&cases)

	docs = append(docs, schema.Document{
		PageContent: fmt.Sprintf("Total number of cases in the system: %d", len(cases)),
		Metadata:    map[string]any{"source": "db:cases:summary"},
	})

	openCount, inProgressCount, closedCount := 0, 0, 0
	for _, c := range cases {
		switch c.Status {
		case models.OPEN:
			openCount++
		case models.IN_PROGRESS:
			inProgressCount++
		case models.CLOSED:
			closedCount++
		}
	}
	docs = append(docs, schema.Document{
		PageContent: fmt.Sprintf(
			"Case status summary: Open=%d, In Progress=%d, Closed=%d",
			openCount, inProgressCount, closedCount,
		),
		Metadata: map[string]any{"source": "db:cases:status_summary"},
	})

	for _, c := range cases {
		caseNumber := safeStr(c.Case_Number)
		title := safeStr(c.Case_Title)
		desc := safeStr(c.Case_Description)
		location := safeStr(c.Location)
		id := safeStr(c.ID)
		incidentDate := time.Time(c.Incident_Date).Format("2006-01-02")

		// Kumpulkan nama detektif yang assigned
		detectiveNames := ""
		for i, d := range c.Detective {
			if i > 0 {
				detectiveNames += ", "
			}
			detectiveNames += safeStr(d.Name)
		}
		if detectiveNames == "" {
			detectiveNames = "None"
		}

		// Kumpulkan nama suspects yang terkait
		suspectNames := ""
		for i, s := range c.Suspects {
			if i > 0 {
				suspectNames += ", "
			}
			suspectNames += safeStr(s.Full_Name)
		}
		if suspectNames == "" {
			suspectNames = "None"
		}

		docs = append(docs, schema.Document{
			PageContent: fmt.Sprintf(
				"Case ID: %s | Case Number: %s | Title: %s | Description: %s | Status: %s | Location: %s | Incident Date: %s | Assigned Detectives: %s | Suspects: %s",
				id, caseNumber, title, desc, string(c.Status), location, incidentDate, detectiveNames, suspectNames,
			),
			Metadata: map[string]any{"source": fmt.Sprintf("db:case:%s", id)},
		})
	}

	// ==================== DETECTIVES ====================
	var detectives []models.Detective
	db.Preload("Cases").Find(&detectives)

	docs = append(docs, schema.Document{
		PageContent: fmt.Sprintf("Total number of detectives in the system: %d", len(detectives)),
		Metadata:    map[string]any{"source": "db:detectives:summary"},
	})

	for _, d := range detectives {
		id := safeStr(d.ID)
		name := safeStr(d.Name)
		badge := safeStr(d.Badge_Number)
		department := safeStr(d.Department)
		station := safeStr(d.Station)
		phone := safeStr(d.Phone)

		assignedCases := ""
		for i, c := range d.Cases {
			if i > 0 {
				assignedCases += ", "
			}
			assignedCases += safeStr(c.Case_Title)
		}
		if assignedCases == "" {
			assignedCases = "None"
		}

		docs = append(docs, schema.Document{
			PageContent: fmt.Sprintf(
				"Detective ID: %s | Name: %s | Badge Number: %s | Department: %s | Station: %s | Phone: %s | Investigation Style: %s | Assigned Cases: %s",
				id, name, badge, department, station, phone, string(d.Investigation_Style), assignedCases,
			),
			Metadata: map[string]any{"source": fmt.Sprintf("db:detective:%s", id)},
		})
	}

	// ==================== SUSPECTS ====================
	var suspects []models.Suspects
	db.Find(&suspects)

	docs = append(docs, schema.Document{
		PageContent: fmt.Sprintf("Total number of suspects in the system: %d", len(suspects)),
		Metadata:    map[string]any{"source": "db:suspects:summary"},
	})

	// Status summary suspects
	statusCount := map[models.Status_Suspect]int{}
	for _, s := range suspects {
		statusCount[s.Status]++
	}
	docs = append(docs, schema.Document{
		PageContent: fmt.Sprintf(
			"Suspect status summary: Arrested=%d, Released=%d, Wanted=%d, Under Investigation=%d, Eyewitness=%d",
			statusCount[models.Arrested],
			statusCount[models.Released],
			statusCount[models.Wanted],
			statusCount[models.Under_Investigation],
			statusCount[models.Eyewitness],
		),
		Metadata: map[string]any{"source": "db:suspects:status_summary"},
	})

	for _, s := range suspects {
		id := safeStr(s.ID)
		caseID := safeStr(s.Case_ID)
		idCard := safeStr(s.ID_card_Number)
		name := safeStr(s.Full_Name)
		address := safeStr(s.Address)
		phone := safeStr(s.Phone)
		occupation := safeStr(s.Occupation)
		alibi := safeStr(s.Alibi)
		dob := time.Time(s.Date_Of_Birth).Format("2006-01-02")

		docs = append(docs, schema.Document{
			PageContent: fmt.Sprintf(
				"Suspect ID: %s | Case ID: %s | ID Card: %s | Name: %s | Gender: %s | Date of Birth: %s | Address: %s | Phone: %s | Occupation: %s | Alibi: %s | Status: %s",
				id, caseID, idCard, name, string(s.Gender), dob, address, phone, occupation, alibi, string(s.Status),
			),
			Metadata: map[string]any{"source": fmt.Sprintf("db:suspect:%s", id)},
		})
	}

	// ==================== SIMPAN KE VECTOR STORE ====================
	_, err := v.Store.AddDocuments(ctx, docs)
	if err != nil {
		return fmt.Errorf("failed to embed DB snapshot: %v", err)
	}

	fmt.Println("✅ Database snapshot embedded to vector store!")
	return nil
}

// safeStr dereference pointer string, return empty string jika nil
func safeStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
