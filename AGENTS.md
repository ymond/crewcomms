# Crew Comms Agent Roles

Crew Comms uses a strict 4-agent, test-first development cycle:

1. Planner → Defines the next increment (no code, no tests).
2. Tester → Writes failing tests only.
3. Implementer → Writes the minimum code to make tests pass.
4. Chronicler → Updates documentation & .context logs.

Then the cycle repeats:
Planner → Tester → Implementer → Chronicler → Planner → …

## Rules

- No agent may perform another agent's responsibilities.
- Planner: plans only.
- Tester: tests only.
- Implementer: implementation only.
- Chronicler: docs/context only.
- Every PR must be authored by a different agent from the previous PR.
- Only one behavior per PR (single scope).
- Always begin a session by reading `.context/HANDOFF.md`; it contains the BLEER startup
  ritual and will cascade into the latest session log and persona playbook.
