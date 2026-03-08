# Minimum Spec (Phase 1)

> Goal: Capture the *minimum* needed to begin without back-and-forth.

## Project
- Name: PIOS Destruction Simulator
- One-liner: A web-based physics simulator where users click to spawn cannonballs that crash into a tall stack of glass blocks.
- Primary user: PIOS Framework Evaluators
- Primary goal: Provide a visually impressive, zero-compile demo of the PIOS execution contract.
- Non-goals (what we will NOT do in v1): Complex games, scoring, textures, mobile optimization.

## Constraints
- Time constraint: Fast (under 10 minutes)
- Must-use tech (if any): HTML5 Canvas, Matter.js (via CDN)
- Must-avoid tech (if any): Build steps (Webpack/Vite), backend servers.

## Success
- MVP success metric: Clicking spawns a ball that knocks over a tower of blocks.
- Failure condition: Unresponsive physics or requirement to compile code.

## Tooling Context
- Expected repo type: Static Webpage
- Deployment target: Local browser

## Next Step
- Move to `spec-lock.md` and resolve only the highest-impact unknowns.
