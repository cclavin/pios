# Plan Lock (Phase 3)

> Goal: Convert Spec Lock into architecture + execution plan.

## Architecture Overview
- Components: A single `index.html` file containing scoped CSS and vanilla Javascript.
- Data flow: User clicks -> JS captures coordinates -> Matter.js spawns a circular body -> Physics engine ticks -> Canvas renders bodies.
- External dependencies: Matter.js loaded via a CDN (`https://cdnjs.cloudflare.com/ajax/libs/matter-js/0.19.0/matter.min.js`).

## Interfaces
- Input: Mouse clicks on the canvas.
- Output: 2D Canvas rendering of rigid bodies.

## Environments
- Local dev: Double clicking `index.html` in any modern web browser.

## Testing Strategy
- Smoke checks: Open file, visually confirm tower is drawn, click to confirm blocks fall.

## Task Breakdown Preview
- Milestone 1: Setup HTML boundaries and engine.
- Milestone 2: Build the glass block tower.
- Milestone 3: Implement click interactions for projectiles.

## Exit Criteria
- [x] Architecture is coherent
- [x] Testing strategy exists
- [x] A first milestone can be scaffolded immediately
