# Minimum Spec (Phase 1)

> Goal: Capture the *minimum* needed to begin without back-and-forth.

## Project
- Name: PIOS Excuse Generator Widget
- One-liner: A customizable, CWV-friendly embeddable widget that generates random excuses.
- Primary user: Website owners, e-commerce stores wanting to add a fun element to their 404 pages or footers.
- Primary goal: Provide a practical, real-world example of an embeddable JS widget built using PIOS.
- Non-goals: Backend API, persistent state, complex UI frameworks (React/Vue).

## Constraints
- Time constraint: Fast (under 20 minutes)
- Must-use tech: Vanilla JS, scoped CSS, IIFE structure.
- Must-avoid tech: External libraries that bloat load time, heavy frameworks.

## Success
- MVP success metric: Adding a `<script>` tag and a `<div id="pios-excuse-widget">` to an HTML page correctly renders a functional button that updates text.
- Failure condition: Widget causes CLS (Cumulative Layout Shift) issues or bleeds CSS into the parent page.

## Tooling Context
- Expected repo type: Embeddable Script
- Deployment target: Local browser test

## Next Step
- Move to `spec-lock.md` and resolve only the highest-impact unknowns.
