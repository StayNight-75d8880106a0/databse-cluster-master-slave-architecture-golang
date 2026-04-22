# ALL.md

## Complete System Flow, Logic, and Usage (Crime App Management)

### 1. Overview
This application is a criminal case management system with four main modules: Dashboard, Cases, Detectives, and Suspects. All features are accessible via the sidebar navigation. The UI is modern, responsive, and provides clear feedback for all actions.

### 2. User Journey (Step-by-Step)

#### a. Dashboard
When the application is opened, the user lands on the Dashboard.
The dashboard displays key statistics (total cases, open/closed cases, detectives) and a pie chart of case statuses.
Recent cases are listed for quick access.

#### b. Case Management
- User navigates to "Cases" from the sidebar.
- The Cases page lists all cases with search, filter, and pagination.
- User can create a new case, view details, edit, or delete cases.
- Clicking a case opens the Case Detail page, showing all info and related suspects.

#### c. Detective Management
- User navigates to "Detectives" from the sidebar.
- The Detectives page lists all detectives with search, filter, and pagination.
- User can add a new detective, view details, edit, or delete detectives.
- Clicking a detective opens the Detective Detail page, showing all info and assigned cases.

#### d. Suspect Management
- Suspects are managed within the Case Detail page, in the "Suspects" section.
- User can add, edit, or delete suspects for a case.
- All suspect CRUD operations are performed in a modal or inline form.

### 3. CRUD Operations Summary
- **Create:** Accessible via "+" buttons or "Add" actions. Opens a form/modal for input. Validation errors are shown inline.
- **Read:** All lists are paginated and searchable. Detail pages show full entity info.
- **Update:** Edit actions open forms with pre-filled data. Changes are saved on submit.
- **Delete:** Delete actions prompt for confirmation before removal.

### 4. Error Handling
- All forms validate required fields before submission.
- Errors (network, validation, etc.) are shown as banners or below fields.
- Loading states are shown with spinners or skeletons.

### 5. UI Consistency
- All pages use consistent layouts, icons, and feedback.
- Navigation is always available via the sidebar.
- Success and error messages are clear and user-friendly.

### 6. End-to-End Flow Example
1. User opens the application and sees the Dashboard.
2. User navigates to "Cases" and creates a new case.
3. User opens the new case, adds suspects, and assigns detectives.
4. User edits or deletes cases, detectives, or suspects as needed.
5. User returns to the Dashboard to see updated statistics.

### 7. Notes
- All instructions above are based on the current UI and reflect the actual user experience.
- The system is designed for clarity, efficiency, and ease of use for investigative workflows.
