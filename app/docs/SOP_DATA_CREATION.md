# SOP_DATA_CREATION.md

## Standard Operating Procedure (SOP): Sequential Data Creation Steps (Crime App Management)

This SOP outlines the recommended order and steps for creating data in the Crime App system to ensure data integrity and smooth workflow.

---

### 1. Create Detectives
- Navigate to the "Detectives" page.
- Add all detectives who will be assigned to cases.
- Ensure each detective has a unique badge number and complete information.

### 2. Create Cases
- Go to the "Cases" page.
- Create new cases and assign one or more detectives to each case.
- Fill in all required case details (title, description, incident date, location, status).

### 3. Add Suspects to Cases
- Open the detail page for a specific case.
- In the "Suspects" section, add suspects related to the case.
- Enter all required suspect information (ID Card Number, name, gender, etc.).

### 4. (Optional) Update or Edit Data
- If needed, update detective, case, or suspect information via their respective detail pages.
- Use the edit (pencil) icon to modify data.

### 5. (Optional) Delete Data
- If a record is no longer needed, use the delete (trash) icon on the relevant page.
- Confirm deletion in the dialog prompt.

---

## Recommended Sequence
1. **Detective** → 2. **Case** (assign detective) → 3. **Suspect** (link to case)

- Always create detectives first so they can be assigned to cases.
- Always create cases before adding suspects, as suspects must be linked to a case.

---

## Notes
- Follow this order to avoid errors such as missing assignments or unlinked data.
- All data creation, update, and deletion should be performed via the UI for consistency.
- Validate all required fields before submitting forms.

---

This SOP ensures a logical, error-free workflow for data entry in the system.
