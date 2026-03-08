# Plan Lock (Phase 3)

> Goal: Convert Spec Lock into architecture + execution plan.

## Architecture Overview
- Components: A single `index.html` file housing inline styles and a primary `script.js` file housing the game states.
- Data flow: Global `gameState` variable ('menu', 'cat', 'platformer'). A main Render loop delegates drawing based on state. 
- Input Handling: Global event listeners for `mousedown`, `keydown`, and `keyup` that route data to the active game state.

## Interfaces
- Menu State: Drawn via HTML DOM overlay to easily manage clicks.
- Game States: Drawn purely via HTML5 Canvas context.

## State 1: Sleeping Cat
- Logic: Draws a visually distinct curled-up cat on the canvas. Tracks mouse clicks or key presses to trigger a "wake up/stretch" animation or visual change.

## State 2: Ice Platformer
- Logic: 
  - Player object (a bright glowing square).
  - Physics implementation: Velocity mapping (vx, vy), gravity, and heavily reduced friction (ice physics) to build momentum.
  - Mechanic: Holding "down/S" increases gravity effect/speed on slopes, allowing momentum build up to hit a "big jump".
  - Level Design: A hardcoded array of platform objects forming a downhill slope leading into a massive ramp.

## Task Breakdown Preview
- Milestone 1: Scaffold Menu and Canvas rendering loop.
- Milestone 2: Implement Sleeping Cat interactions.
- Milestone 3: Implement Ice Platformer physics and level.

## Exit Criteria
- [x] Architecture is coherent
- [x] Testing strategy exists
- [x] A first milestone can be scaffolded immediately
