---
name: notes
version: 0.1.0
author: goclaw
description: "Quick notes, lists, and ideas — stored locally as markdown files"
category: builtin
tags: [notes, memo, lists, ideas, journal, quick-capture]
---
# Notes

You can save, search, read, and manage notes for the user. Notes are stored as markdown files in the `~/.goclaw/notes/` directory.

## Creating notes

```bash
# Ensure directory exists
mkdir -p ~/.goclaw/notes

# Quick note (timestamped)
cat > ~/.goclaw/notes/$(date +%Y%m%d-%H%M%S)-note.md << 'EOF'
# Quick note

Content of the note here.

Tags: #idea #work
EOF

# Named note
cat > ~/.goclaw/notes/shopping-list.md << 'EOF'
# Shopping List

- [ ] Leite
- [ ] Pão
- [ ] Ovos
- [ ] Frutas
EOF

# Append to existing note
echo "- [ ] Café" >> ~/.goclaw/notes/shopping-list.md

# Journal entry (daily)
cat >> ~/.goclaw/notes/journal-$(date +%Y-%m-%d).md << EOF

## $(date +%H:%M) — Entry

What happened today...

EOF
```

## Reading notes

```bash
# List all notes
ls -lt ~/.goclaw/notes/ | head -20

# Read a specific note
cat ~/.goclaw/notes/shopping-list.md

# Search across all notes
grep -rl "SEARCH_TERM" ~/.goclaw/notes/ 2>/dev/null
grep -n "SEARCH_TERM" ~/.goclaw/notes/*.md 2>/dev/null

# Recent notes (last 7 days)
find ~/.goclaw/notes/ -name "*.md" -mtime -7 -exec ls -lt {} +
```

## Editing notes

```bash
# Replace content
cat > ~/.goclaw/notes/FILENAME.md << 'EOF'
# Updated title

New content here.
EOF

# Mark todo as done (replace "- [ ]" with "- [x]")
sed -i 's/- \[ \] Leite/- [x] Leite/' ~/.goclaw/notes/shopping-list.md
```

## Deleting notes

```bash
# Delete a note (always confirm first!)
rm ~/.goclaw/notes/FILENAME.md

# Archive old notes
mkdir -p ~/.goclaw/notes/archive
mv ~/.goclaw/notes/old-note.md ~/.goclaw/notes/archive/
```

## Note types

| Type | Naming convention | Example |
|------|------------------|---------|
| Quick note | `YYYYMMDD-HHMMSS-note.md` | `20260214-153000-note.md` |
| Named note | `descriptive-name.md` | `shopping-list.md` |
| Daily journal | `journal-YYYY-MM-DD.md` | `journal-2026-02-14.md` |
| Project note | `project-NAME.md` | `project-goclaw.md` |
| Todo list | `todo-NAME.md` | `todo-this-week.md` |

## Tips

- Always confirm before deleting notes.
- Use descriptive names so notes are easy to find later.
- For todo lists, use `- [ ]` / `- [x]` markdown checkboxes.
- Use tags at the bottom of notes (e.g., `Tags: #work #urgent`) for easier searching.
- When the user says "anota isso" or "salva isso", create a quick note with timestamp.
- Read the note back to the user after creating it for confirmation.
- For large collections, suggest organizing by folders/categories.

## Triggers

note, save this, anota isso, salva isso, create a note, my notes,
lista de compras, shopping list, todo list, journal, diário, lembrar disso
