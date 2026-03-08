---
pios_contract_version: "0.4"
---
# TASKS (Phase 4)

> Rule: tasks must be small, testable, and include acceptance criteria.

## Milestone 1 - Benchmark Execution

### [x] TASK-001: Scaffold Menu UI and Game Loop
- **Status:** Done
- **Dependencies:** None
- **Acceptance Criteria:**
  - [x] `index.html` exists with an overlay DOM menu and a fullscreen `<canvas>`.
  - [x] The menu has two distinct buttons: "Sleeping Cat" and "Ice Platformer".
  - [x] A core JavaScript `requestAnimationFrame` loop exists that routes drawing logic based on the selected state.
  - [x] An escape/back button exists to return to the menu from a game.

### [x] TASK-002: Implement Sleeping Cat Toy
- **Status:** Done
- **Dependencies:** TASK-001
- **Acceptance Criteria:**
  - [x] Drawn on HTML5 canvas. Visualizes a sleeping cat (using basic Canvas shapes to keep it zero-dependency).
  - [x] Clicking on or near the cat, or pressing keys causes the cat to visibly react (e.g. open an eye, shift position).
  - [x] The cat returns to a sleeping state shortly after interaction ceases.

### [x] TASK-003: Implement Ice Platformer Simulator
- **Status:** Done
- **Dependencies:** TASK-001
- **Acceptance Criteria:**
  - [x] Player can use WASD/Arrows to move a character box left/right and jump.
  - [x] Uses "ice physics" (low friction) allowing momentum buildup.
  - [x] Implements a "tuck down" mechanic (holding S/Down arrow) that accelerates the player character downhill.
  - [x] Level design includes a significant downhill slope and a corresponding "big jump" ramp to test the mechanics.
