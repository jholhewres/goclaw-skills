---
name: claude-code
description: "Full-stack coding assistant powered by Claude Code CLI"
---
# Claude Code

Use the **bash** tool with the Claude Code CLI for advanced coding tasks.

## Setup

1. **Check if installed:**
   ```bash
   command -v claude && claude --version
   ```

2. **Install:**
   ```bash
   npm install -g @anthropic-ai/claude-code
   ```

3. **Auth:** Use the vault for `ANTHROPIC_API_KEY`. Stored keys are auto-injected as env vars (UPPERCASE).
   ```bash
   # Save to vault (key name lowercase)
   vault_save anthropic_api_key "sk-ant-..."

   # Or interactive
   claude setup-token
   ```

## Usage
```bash
claude -p "fix the authentication bug in auth.ts" --allowedTools bash,read,write
claude -p "review this code for security issues" --permission-mode plan
```

## Tips
- Be specific in prompts
- For read-only analysis, use --permission-mode plan
- Check auth: claude status
