# Development Plan

Crew Comms uses the cycle:
Planner → Tester → Implementer → Chronicler → Planner → ...

This scaffold PR = Planner.

Coordination follows the BLEER protocol documented in `.context/HANDOFF.md`; every session log
records the acting persona and next handoff target.

## Upcoming Milestones (each a future Planner PR)

- Wire Cobra root (no commands)
- Add `instance ls` command scaffolding
- Add TOML schema and read/write (tests first)
- Create webhook handler skeleton & TwiML encoder
- Implement slash command parsing
- Implement membership logic
- Broadcast formatting
- Integrate Twilio send pipeline
- Add 1Password sync
- Add Bubble Tea TUI (later)
