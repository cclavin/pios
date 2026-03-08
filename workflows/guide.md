# PIOS Operating Guide

This document captures the sequential flows to progress a project from an empty folder to a finished MVP using the PIOS [execution contract](../docs/positioning.md).

---

## 1. Whiteboard Flow
*Objective: Brainstorm and commit to a scope.*
1) Fill `templates/min-spec.md`
2) Convert to `templates/spec-lock.md`
3) Produce `templates/plan-lock.md`
4) Draft `templates/tasks.md`
5) Start scaffolding

---

## 2. Scaffold Flow
*Objective: Turn the plan into a bootable repository.*
1) Read Plan Lock + Tasks
2) Generate file tree
3) Add bootstrap scripts / Makefile
4) Make repo runnable quickly (even if stubbed)
5) Iterate via autopilot loop (implement → test → fix → doc → summarize)

---

## 3. Audit Flow
*Objective: Review for debt, security, and simplification.*
1) Check for conflicting standards
2) Identify top risks
3) Identify simplifications
4) Verify tests exist or propose minimal set
5) Update Decision Log

---

## 4. Launch Flow
*Objective: Define the path to production.*
1) Define minimal deployment target
2) Ensure secrets handling is correct
3) Add basic observability
4) Confirm rollback plan (even if manual)
5) Write a launch checklist
