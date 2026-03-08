# Plan Lock (Phase 3)

> Goal: Convert Spec Lock into architecture + execution plan.

## Architecture Overview
- Components: A single `widget.js` file and an `index.html` demo file.
- Data flow: User clicks "Generate" button -> JS selects random string from array -> JS updates DOM element text.
- Architecture: Immediately Invoked Function Expression (IIFE) to avoid polluting global scope.
- Injection Strategy: The script looks for `<div id="pios-excuse-widget">` and injects scoped HTML/CSS inside it. Uses data attributes (`data-button-text`) for easy customization.

## Interfaces
- Input: HTML data attributes on the container `div`.
- Output: DOM injection of the widget UI.

## Testing Strategy
- Smoke checks: Open `index.html`, verify the widget loads without console errors, and verify the excuse changes on click.
- CWV checks: Verify height is pre-allocated (to prevent CLS).

## Task Breakdown Preview
- Milestone 1: Setup file structure and widget loader.
- Milestone 2: Implement logic, data, and styling.
- Milestone 3: Demo the implementation.

## Exit Criteria
- [x] Architecture is coherent
- [x] Testing strategy exists
- [x] A first milestone can be scaffolded immediately
