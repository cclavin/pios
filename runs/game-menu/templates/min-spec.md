# Minimum Spec (Phase 1)

> Goal: Capture the *minimum* needed to begin without back-and-forth.

## Project
- Name: PIOS Dual-Game Menu Benchmark
- One-liner: A game menu interface allowing selection between a serene "Sleeping Cat" toy and an active "Ice Platformer" game.
- Primary user: PIOS Framework Evaluators
- Primary goal: Prove the PIOS execution contract can generate cohesive multi-state UI logic and complex canvas physics (ice momentum mechanics) in a single run.
- Non-goals (what we will NOT do in v1): Audio, complex game over states, external asset loading (we will draw everything on canvas).

## Constraints
- Time constraint: Fast (<20 minutes)
- Must-use tech: Vanilla JS, HTML5 Canvas.
- Must-avoid tech: Heavy external game engines (Phaser/Unity), build steps.

## Success
- MVP success metric: Opening `index.html` shows a menu. Clicking Game 1 shows a sleeping cat interacting with clicks. Clicking Game 2 shows a playable 2D platformer with ice physics and a jump mechanic.
- Failure condition: Games fail to load, or mechanics are broken.

## Tooling Context
- Expected repo type: Static Webpage
- Deployment target: Local browser

## Next Step
- Move to `spec-lock.md` and resolve only the highest-impact unknowns.
