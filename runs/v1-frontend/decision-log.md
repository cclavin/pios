# Decision Log

## Entries

- Date: 2026-03-08
  Decision: Use next-themes for dark mode
  Options: Manual class toggle, CSS variables only, next-themes
  Why: Handles SSR flash via suppressHydrationWarning; localStorage persistence built-in
  Consequences: Requires `"use client"` ThemeProvider wrapper in App Router layout

- Date: 2026-03-08
  Decision: Glassmorphism via Tailwind backdrop-blur + custom .glass CSS classes
  Options: CSS-in-JS, custom CSS module, Tailwind utilities only
  Why: Tailwind utilities cover most cases; custom .glass/.glass-light classes handle dark/light variants cleanly
  Consequences: backdrop-blur not supported in very old browsers (acceptable for this use case)

- Date: 2026-03-08
  Decision: Scaffold in-repo with mock data only (no API calls)
  Options: Mock data, MSW, real API
  Why: Spec explicitly scoped to UI scaffold only; avoids backend dependency in v1
  Consequences: Must wire up real API in v2; NEXT_PUBLIC_API_URL kept in .env.example for future use

- Date: 2026-03-08
  Decision: create-next-app into sibling dir then copy into PIOS project dir
  Options: Manual scaffold, copy approach
  Why: create-next-app refuses to run in non-empty directories; copying is safe and idempotent
  Consequences: pios-app sibling dir can be deleted after copy
