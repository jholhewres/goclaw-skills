---
name: calendar
version: 0.2.0
author: goclaw
description: "Google Calendar integration via gcalcli or gog CLI"
category: productivity
tags: [google, calendar, scheduling, events, meetings]
requires:
  any_bins: [gcalcli, gog]
---
# Calendar

You can manage Google Calendar events. Requires `gcalcli` or `gog` CLI.

## Using gcalcli

```bash
# List upcoming events
gcalcli agenda --nostarted --details location --details description

# List events for a specific date range
gcalcli agenda "2026-02-14" "2026-02-21"

# Create an event
gcalcli add --title "Meeting" --where "Room 1" --when "tomorrow 3pm" --duration 60 --description "Weekly sync"

# Quick add (natural language)
gcalcli quick "Meeting with João tomorrow at 3pm for 1 hour"

# Delete an event
gcalcli delete --query "Meeting"

# Calendar for the month
gcalcli calw 4
```

## Using gog (Google Workspace CLI)

```bash
# List upcoming events
gog calendar list --days 7

# Create event
gog calendar create --title "Meeting" --start "2026-02-15T15:00:00" --duration "1h" --description "Weekly sync"

# Delete event
gog calendar delete --id EVENT_ID
```

## Tips

- Always confirm with the user before creating or deleting events.
- Parse natural language dates: "tomorrow", "next Monday", "in 2 hours".
- Use the user's timezone (check USER.md or config).
- For recurring events, mention the recurrence pattern.
- When listing events, show: title, date/time, duration, location.
- If neither `gcalcli` nor `gog` is installed, suggest installation.

## Triggers

what's on my calendar, schedule a meeting, create an event,
check my schedule, minha agenda, agendar, marcar reunião, calendario
