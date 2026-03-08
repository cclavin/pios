---
pios_contract_version: "0.4"
---
# TASKS

## Milestone 1 - Scaffold and Bootstrap

### [x] TASK-001: Initialize Next.js + Tailwind project
- **Status:** Done
- **Dependencies:** None
- **Acceptance Criteria:**
  - [x] `npm run dev` starts the app
  - [x] Tailwind styles are applied

### [x] TASK-002: Add next-themes for dark mode
- **Status:** Done
- **Dependencies:** TASK-001
- **Acceptance Criteria:**
  - [x] ThemeProvider wraps the app
  - [x] Dark/light toggle works without flash

### [x] TASK-003: Add .env.example
- **Status:** Done
- **Dependencies:** None
- **Acceptance Criteria:**
  - [x] `.env.example` exists
  - [x] No real secrets committed

---

## Milestone 2 - Core UI Components

### [x] TASK-004: Create mock data module
- **Status:** Done
- **Dependencies:** TASK-001
- **Acceptance Criteria:**
  - [x] `lib/mockData.ts` exports tasks array

### [x] TASK-005: Build TaskCard component
- **Status:** Done
- **Dependencies:** TASK-004
- **Acceptance Criteria:**
  - [x] Glassmorphism card with title, status badge, priority

### [x] TASK-006: Build TaskGrid component
- **Status:** Done
- **Dependencies:** TASK-005
- **Acceptance Criteria:**
  - [x] Responsive 3-column grid of TaskCards

### [x] TASK-007: Build Navbar with dark mode toggle
- **Status:** Done
- **Dependencies:** TASK-002
- **Acceptance Criteria:**
  - [x] Logo/title visible
  - [x] Toggle button switches theme

### [x] TASK-008: Build StatsBar summary
- **Status:** Done
- **Dependencies:** TASK-004
- **Acceptance Criteria:**
  - [x] Shows total, complete, pending counts

---

## Milestone 3 - Hardening

### [x] TASK-009: Wire up dashboard page
- **Status:** Done
- **Dependencies:** TASK-006, TASK-007, TASK-008
- **Acceptance Criteria:**
  - [x] `/` renders full dashboard with all components

### [x] TASK-010: Final polish and decision log
- **Status:** Done
- **Dependencies:** All
- **Acceptance Criteria:**
  - [x] STATUS.md updated
  - [x] decision-log.md has entries
