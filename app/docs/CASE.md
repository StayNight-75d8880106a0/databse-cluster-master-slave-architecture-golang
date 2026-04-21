# CASE.md

## Step-by-Step UI Instructions for Case Management (LangChain AI Training)

### 1. Viewing All Cases (Read)
- Navigate to the "Cases" page from the sidebar.
- The page displays a paginated table/list of cases with columns: Case Number, Title, Description, Status, Incident Date, Assigned Detectives, and Actions.
- Use the search bar to filter cases by title, number, or description.
- Use the status filter dropdown to filter by case status (Open, In Progress, Closed).
- Use pagination controls to navigate between pages.
- Each row has action buttons: View (eye icon), Edit (pencil icon), Delete (trash icon).

### 2. Creating a New Case (Create)
- Click the "+ New Case" button (Plus icon) at the top of the Cases page.
- You are redirected to the "Create New Case" form.
- Fill in the following fields:
  - Title (required)
  - Description (required)
  - Incident Date (required, date picker)
  - Location (required)
  - Status (dropdown: Open, In Progress, Closed)
  - Assign Detectives (multi-select dropdown)
- Click "Submit" to create the case.
- On success, you are redirected back to the Cases page and see the new case in the list.
- If you click "Cancel", you are returned to the Cases page without saving.
- If validation fails, error messages are shown below the relevant fields.

### 3. Viewing Case Details (Read)
- Click the "View" (eye icon) on a case row.
- The Case Detail page displays all case information, including:
  - Title, Description, Incident Date, Location, Status
  - List of assigned detectives
  - List of suspects (with option to manage suspects)
- You can return to the Cases page using the "Back" button.

### 4. Editing a Case (Update)
- On the Case Detail page, click the "Edit" (pencil icon) button.
- The form fields become editable.
- Update any field as needed.
- Click "Save" to update the case.
- On success, the updated information is shown.
- Click "Cancel" to discard changes and revert to view mode.
- If validation fails, error messages are shown below the relevant fields.

### 5. Deleting a Case (Delete)
- On the Cases page, click the "Delete" (trash icon) button for a case.
- A confirmation dialog appears: "Are you sure you want to delete this case?"
- Click "Confirm" to delete, or "Cancel" to abort.
- On success, the case is removed from the list.

### 6. Error Handling
- If an operation fails (e.g., network error, validation error), an error message is displayed at the top or below the relevant field.
- Required fields are validated before submission.

### 7. Notes
- All actions are performed via the UI and reflect changes immediately after success.
- The UI uses clear labels, icons, and feedback for each action.
- All CRUD operations are accessible from the Cases page or Case Detail page.
