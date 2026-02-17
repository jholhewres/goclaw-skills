---
name: github
version: 0.3.0
author: devclaw
description: "Full GitHub integration via gh CLI — issues, PRs, releases, CI, search"
category: development
tags: [github, git, issues, pull-requests, repositories, ci, releases]
requires:
  bins: [gh]
---
# GitHub

You have full access to GitHub using the `gh` CLI. Make sure `gh auth status` shows authenticated.

## Setup

1. **Check if installed:**
   ```bash
   command -v gh && gh --version
   ```

2. **Install:**
   ```bash
   # macOS
   brew install gh

   # Ubuntu / Debian (see https://github.com/cli/cli/blob/trunk/docs/install_linux.md)
   type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y)
   curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
   sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
   echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
   sudo apt update && sudo apt install -y gh
   ```

3. **Auth:**
   ```bash
   # Interactive (browser)
   gh auth login

   # Non-interactive: store token in vault; auto-injects as GITHUB_TOKEN (UPPERCASE)
   vault_save github_token "ghp_..."
   echo $GITHUB_TOKEN | gh auth login --with-token
   ```

## Issues

```bash
# List issues
gh issue list -R OWNER/REPO --limit 10
gh issue list -R OWNER/REPO --state closed --label "bug"

# View issue details
gh issue view NUMBER -R OWNER/REPO

# Create issue
gh issue create -R OWNER/REPO --title "TITLE" --body "BODY" --label "bug" --assignee "@me"

# Close issue
gh issue close NUMBER -R OWNER/REPO

# Comment on issue
gh issue comment NUMBER -R OWNER/REPO --body "Comment text"

# Structured output
gh issue list -R OWNER/REPO --json number,title,state,labels,assignees | jq '.'
```

## Pull Requests

```bash
# List PRs
gh pr list -R OWNER/REPO --limit 10
gh pr list -R OWNER/REPO --state merged

# View PR details
gh pr view NUMBER -R OWNER/REPO

# View diff
gh pr diff NUMBER -R OWNER/REPO

# Check CI status
gh pr checks NUMBER -R OWNER/REPO

# Merge PR (squash, merge, or rebase)
gh pr merge NUMBER -R OWNER/REPO --squash --delete-branch

# Create PR
gh pr create -R OWNER/REPO --title "TITLE" --body "BODY" --base main
```

## CI / GitHub Actions

```bash
# List recent workflow runs
gh run list -R OWNER/REPO --limit 5
gh run list -R OWNER/REPO --workflow "ci.yml"

# View run details and logs
gh run view RUN_ID -R OWNER/REPO
gh run view RUN_ID -R OWNER/REPO --log-failed

# Re-run a failed workflow
gh run rerun RUN_ID -R OWNER/REPO --failed
```

## Releases

```bash
# List releases
gh release list -R OWNER/REPO --limit 5

# View latest release
gh release view --latest -R OWNER/REPO

# Create release
gh release create TAG -R OWNER/REPO --title "TITLE" --notes "Release notes" --generate-notes

# Upload assets to release
gh release upload TAG ./file.zip -R OWNER/REPO
```

## Search & API

```bash
# Search repos
gh search repos "QUERY" --limit 10

# Search code
gh search code "QUERY" -R OWNER/REPO --limit 10

# Raw API call (for anything not covered)
gh api repos/OWNER/REPO/commits --paginate | jq '.[0:5] | .[].commit.message'
gh api repos/OWNER/REPO/contributors | jq '.[] | {login, contributions}'
```

## Repo management

```bash
# Clone, fork
gh repo clone OWNER/REPO
gh repo fork OWNER/REPO

# View repo info
gh repo view OWNER/REPO

# Create repo
gh repo create NAME --public --description "DESC" --clone
```

## Tips

- Always use `-R OWNER/REPO` to target the correct repository.
- Use `--json` flag for structured output, then filter with `jq`.
- Use `--limit` to cap results and avoid overwhelming output.
- Check `gh auth status` if you get permission errors.
- For large diffs, use `gh pr diff NUMBER | head -100`.

## Triggers

github, issue, pull request, PR, workflow, CI, release, create issue,
check repo, merge PR, check CI
