# Plan Lock (Phase 3)

## Architecture Overview
- Components: ThemeProvider, Navbar, TaskCard, TaskGrid, StatsBar, mock data module
- Data flow: mockTasks array → TaskGrid → TaskCard (no fetch in v1)
- External dependencies: next-themes, tailwindcss, @tailwindcss/typography (optional)
- State & persistence: Theme stored in localStorage via next-themes

## Interfaces
- API surface: None (mock data only in v1)
- Auth model: None
- Error handling strategy: None required (static mock)

## Environments
- Local dev: `npm run dev` → http://localhost:3000
- CI: None in v1
- Staging: None
- Prod: Vercel (future)

## Testing Strategy
- Unit: None in v1 (UI scaffold only)
- Integration: None
- E2E: None
- Smoke checks: Visual inspection in browser

## Observability
- Logging: None
- Metrics: None
- Tracing: None

## Security
- Secrets handling: .env.example provided; no real secrets
- Access control: None
- Threat model notes: Public static UI; no threats in v1

## Cost / Scaling Notes
- Expected load: Local dev only
- Bottlenecks: None
- Scaling strategy: N/A

## Task Breakdown Preview
- Milestone 1: Bootstrap Next.js + Tailwind + next-themes
- Milestone 2: Build TaskCard, TaskGrid, Navbar, StatsBar with mock data
- Milestone 3: Polish, .env.example, README

## Exit Criteria
- [x] Architecture is coherent
- [x] Testing strategy exists
- [x] A first milestone can be scaffolded immediately

## Next Step
- Convert milestones into `TASKS.md` with acceptance criteria.
