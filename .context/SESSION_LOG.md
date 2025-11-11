# Session Log Index

Crew Comms tracks every agent session in timestamped files within `./.context/session_logs/`.
The session ledger is the authoritative source for persona adoption and handoff notes.
It works hand-in-hand with the BLEER protocol documented in `.context/HANDOFF.md`.

## Quickstart for Incoming Agents
1. **Locate Latest Entry** — List files in `.context/session_logs/` and select the one with the
   highest timestamp (lexicographically). That file contains the persona that last acted and the
   instructions for the next persona.
2. **Adopt Persona** — Load the matching playbook from `.context/personas/` and follow its
   checklist.
3. **Execute Role** — Perform responsibilities defined in the playbook only.
4. **Relay via BLEER** — When finished, create a new session file and update this index as
   described below.

## How to Log a New Session
1. Create a new Markdown file in `.context/session_logs/` named `YYYYMMDDThhmmssZ_session-XXX_role.md`.
2. Include sections for Summary, Cycle Position (current role, next role, rotation check), Decisions,
   Blockers, and Notes for Future Agents.
3. Document any reverse handoff request explicitly ("Next Persona" may loop back upstream).
4. Append the new entry to the "Active Sessions" table below without removing previous rows.
5. Reference `.context/HANDOFF.md` if you need to escalate or choose an alternate coordination
   mechanism.

## Active Sessions
| Session | Timestamp (UTC)       | Role       | File |
|---------|-----------------------|------------|------|
| 000     | 2024-01-01T00:00:00Z | Planner    | [20240101T000000Z_session-000_planner.md](session_logs/20240101T000000Z_session-000_planner.md) |
| 001     | 2025-01-15T12:00:00Z | Chronicler | [20250115T120000Z_session-001_chronicler.md](session_logs/20250115T120000Z_session-001_chronicler.md) |

## Dormancy Tracking
Calculate dormancy by comparing the timestamps of the two most recent session files. Record long gaps
in the next session's Notes section to maintain historical awareness of project pauses.

## Historical Files
All archived sessions live in `.context/session_logs/`. Review them chronologically to identify
recurring issues, successful practices, rotation compliance, and state transitions.
