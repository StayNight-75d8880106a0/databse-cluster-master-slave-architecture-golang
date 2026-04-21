# DETECTIVE.md

## Step-by-Step UI Instructions for Detective Management (LangChain AI Training)

### 1. Viewing All Detectives (Read)
- Navigate to the "Detectives" page from the sidebar.
- The page displays a paginated table/list of detectives with columns: Name, Badge Number, Department, Station, Phone, Investigation Style, and Actions.
- Use the search bar to filter detectives by name, badge number, or email.
- Use the department and investigation style filter dropdowns to filter the list.
- Use pagination controls to navigate between pages.
- Each row has action buttons: View (eye icon), Edit (pencil icon), Delete (trash icon).

### 2. Creating a New Detective (Create)
- Click the "+ Add Detective" button (Plus icon) at the top of the Detectives page.
- You are redirected to the "Add New Detective" form.
- Fill in the following fields:
  - Name (required)
  - Badge Number (required)
  - Department (required, dropdown)
  - Station (required)
  - Phone (required)
  - Investigation Style (required, dropdown)
- Click "Submit" to create the detective.
- On success, you are redirected back to the Detectives page and see the new detective in the list.
- If you click "Cancel", you are returned to the Detectives page without saving.
- If validation fails, error messages are shown below the relevant fields.

### 3. Viewing Detective Details (Read)
- Click the "View" (eye icon) on a detective row.
- The Detective Detail page displays all detective information, including:
  - Name, Badge Number, Department, Station, Phone, Investigation Style
  - List of assigned cases
- You can return to the Detectives page using the "Back" button.

### 4. Editing a Detective (Update)
- On the Detective Detail page, click the "Edit" (pencil icon) button.
- The form fields become editable.
- Update any field as needed.
- Click "Save" to update the detective.
- On success, the updated information is shown.
- Click "Cancel" to discard changes and revert to view mode.
- If validation fails, error messages are shown below the relevant fields.

### 5. Deleting a Detective (Delete)
- On the Detectives page, click the "Delete" (trash icon) button for a detective.
- A confirmation dialog appears: "Are you sure you want to delete this detective?"
- Click "Confirm" to delete, or "Cancel" to abort.
- On success, the detective is removed from the list.

### 6. Error Handling
- If an operation fails (e.g., network error, validation error), an error message is displayed at the top or below the relevant field.
- Required fields are validated before submission.

### 7. Notes
- All actions are performed via the UI and reflect changes immediately after success.
- The UI uses clear labels, icons, and feedback for each action.
- All CRUD operations are accessible from the Detectives page or Detective Detail page.
