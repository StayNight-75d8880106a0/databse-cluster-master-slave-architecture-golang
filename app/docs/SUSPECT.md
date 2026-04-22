# SUSPECT.md

## Step-by-Step UI Instructions for Suspect Management (Crime App Management)

### 1. Viewing Suspects for a Case (Read)
- On the Case Detail page, scroll to the "Suspects" section.
- The section displays a list/table of suspects associated with the case, showing: NIK/ID Card Number, Name, Gender, Date of Birth, Address, Phone, Occupation, Alibi, Status, and Actions.
- Use pagination controls if there are many suspects.
- Each row has action buttons: Edit (pencil icon), Delete (trash icon).

### 2. Creating a New Suspect (Create)
- In the "Suspects" section, click the "+ Add Suspect" button (Plus icon).
- A modal or form appears for adding a new suspect.
- Fill in the following fields:
  - ID Card Number (required)
  - Name (required)
  - Gender (required, dropdown: Male/Female)
  - Date of Birth (required, date picker)
  - Address (required)
  - Phone (optional)
  - Occupation (optional)
  - Alibi (required)
  - Status (required, dropdown: Arrested, Released, Wanted, Under Investigation, Eyewitness)
- Click "Submit" to create the suspect.
- On success, the suspect is added to the list.
- If you click "Cancel", the form closes without saving.
- If validation fails, error messages are shown below the relevant fields.

### 3. Editing a Suspect (Update)
- In the "Suspects" section, click the "Edit" (pencil icon) button for a suspect.
- A modal or form appears with the suspect's current data pre-filled.
- Update any field as needed.
- Click "Save" to update the suspect.
- On success, the updated information is shown in the list.
- Click "Cancel" to discard changes and close the form.
- If validation fails, error messages are shown below the relevant fields.

### 4. Deleting a Suspect (Delete)
- In the "Suspects" section, click the "Delete" (trash icon) button for a suspect.
- A confirmation dialog appears: "Are you sure you want to delete this suspect?"
- Click "Confirm" to delete, or "Cancel" to abort.
- On success, the suspect is removed from the list.

### 5. Error Handling
- If an operation fails (e.g., network error, validation error), an error message is displayed at the top or below the relevant field.
- Required fields are validated before submission.

### 6. Notes
- All actions are performed via the UI and reflect changes immediately after success.
- The UI uses clear labels, icons, and feedback for each action.
- All CRUD operations for suspects are managed within the Case Detail page, in the Suspects section.
