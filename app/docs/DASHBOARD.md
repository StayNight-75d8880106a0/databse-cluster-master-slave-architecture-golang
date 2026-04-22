# DASHBOARD.md

## Step-by-Step UI Instructions for Dashboard (Crime App Management)

### 1. Viewing Dashboard Overview
- Navigate to the "Dashboard" page from the sidebar (usually the home page).
- The dashboard displays key statistics in stat cards, such as:
  - Total Cases
  - Open Cases
  - Closed Cases
  - Total Detectives
- A pie chart visualizes the distribution of case statuses (Open, Closed, In Progress).
- A list/table of recent cases is shown, sorted by last updated date.
- Each recent case entry displays: Title, Status, Incident Date, and Assigned Detectives.

### 2. Interacting with Dashboard Elements
- Stat cards are for display only and do not have actions.
- The pie chart is for visualization only.
- Click on a recent case entry to navigate to its Case Detail page.

### 3. Error Handling
- If data fails to load (e.g., network error), an error message is displayed.
- Loading spinners or skeletons are shown while data is being fetched.

### 4. Notes
- The dashboard provides a quick overview of the system's current state.
- All data is read-only; no create, update, or delete actions are available from the dashboard.
- The UI uses clear labels, icons, and visual feedback for each element.
