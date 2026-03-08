---
pios_contract_version: "0.4"
---
# TASKS (Phase 4)

> Rule: tasks must be small, testable, and include acceptance criteria.

## Milestone 1 - Benchmark Execution

### [x] TASK-001: Scaffold HTML/CSS and load Matter.js
- **Status:** Done
- **Dependencies:** None
- **Acceptance Criteria:**
  - [x] `index.html` contains basic boilerpate and a dark theme CSS.
  - [x] Matter.js is loaded via CDN script tag.
  - [x] A `<canvas>` element exists.

### [x] TASK-002: Initialize physics world and block tower
- **Status:** Done
- **Dependencies:** TASK-001
- **Acceptance Criteria:**
  - [x] Matter.js Engine, Render, Runner, and World are initialized.
  - [x] Ground and invisible walls are created to contain physics.
  - [x] A tall "tower" of rigid rectangular bodies (representing glass blocks) is stacked.
  - [x] Blocks are styled with a neon or premium aesthetic (e.g. wireframes or custom colors).

### [x] TASK-003: Click interaction for projectiles
- **Status:** Done
- **Dependencies:** TASK-002
- **Acceptance Criteria:**
  - [x] Clicking on the canvas creates a high-mass, circular body at the click coordinates.
  - [x] The generated ball applies force and destroys the block tower on impact.
