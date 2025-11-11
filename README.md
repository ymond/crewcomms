# Crew Comms

Crew Comms is an SMS-based group messaging system built in Go.

## Status: Scaffold Only
This initial PR only sets up the repository structure, documents, and placeholders.
No logic or tests exist yet.

## Development Loop
Crew Comms uses a strict test-first loop:

1. Planner → plan only
2. Tester → failing tests only
3. Implementer → make tests pass
4. Chronicler → document

Then repeat.

### Persona Handoff Protocol
Start every session with the single instruction `read .context/HANDOFF.md`. That file now serves as
the canonical entrypoint: it explains the BLEER ritual, points you to the latest session log, and
tells you exactly which persona playbook inside `.context/personas/` to adopt next. After executing
your role, follow the same instructions to relay the torch for the next agent.

## Roadmap (high-level)
- Wire Cobra root
- Add TOML handling
- Implement webhook skeleton
- Add slash commands (/join, /leave, /who, /help)
- Add broadcast formatting
- Twilio integration
- 1Password sync
- Bubble Tea TUI

## Diagrams (TODO)
```mermaid
%% TODO: add domain model in future PR
```
