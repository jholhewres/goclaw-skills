---
name: discord
version: 0.1.0
author: devclaw
description: "Discord integration — send messages via webhooks and bots"
category: communication
tags: [discord, messaging, gaming, communities, webhooks]
requires:
  bins: [curl, jq]
  env: [DISCORD_WEBHOOK_URL, DISCORD_BOT_TOKEN]
---
# Discord

Interact with Discord using webhooks or the Bot API.

## Setup

1. **Check existing credentials:**
   ```
   vault_get discord_webhook_url
   vault_get discord_bot_token
   ```
   Use webhook URL for simple notifications; use bot token for full Bot API (channels, messages, guilds).

2. **Webhook (easiest):**
   - Go to Server Settings > Integrations > Webhooks in your Discord server
   - Click "New Webhook", name it, select channel
   - Copy the Webhook URL
   - Save to vault:
     ```
     vault_save discord_webhook_url "https://discord.com/api/webhooks/ID/TOKEN"
     ```
   The URL is auto-injected as `$DISCORD_WEBHOOK_URL`.

3. **Bot API (more features):**
   - Go to https://discord.com/developers/applications
   - Create an application, add a bot, copy the Bot Token
   - Save to vault:
     ```
     vault_save discord_bot_token "your-bot-token"
     ```
   The token is auto-injected as `$DISCORD_BOT_TOKEN`.

## Send Messages (Webhook)

```bash
# Simple message
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello from DevClaw!"}'

# With username and avatar override
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Custom message",
    "username": "DevClaw Bot",
    "avatar_url": "https://example.com/avatar.png"
  }'

# With embed (rich formatting)
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "embeds": [{
      "title": "Notification",
      "description": "Something happened!",
      "color": 3447003,
      "fields": [
        {"name": "Status", "value": "Success", "inline": true},
        {"name": "Time", "value": "Now", "inline": true}
      ]
    }]
  }'
```

## Bot API (More Features)

```bash
# Requires DISCORD_BOT_TOKEN (stored in vault)

# Send message to channel
curl -s -X POST "https://discord.com/api/v10/channels/CHANNEL_ID/messages" \
  -H "Authorization: Bot $DISCORD_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello!"}'

# Get channel messages
curl -s "https://discord.com/api/v10/channels/CHANNEL_ID/messages?limit=10" \
  -H "Authorization: Bot $DISCORD_BOT_TOKEN" | jq '.[] | {author: .author.username, content}'

# List guilds (servers)
curl -s "https://discord.com/api/v10/users/@me/guilds" \
  -H "Authorization: Bot $DISCORD_BOT_TOKEN" | jq '.[] | {id, name}'
```

## Embed Examples

```bash
# Status embed with color coding
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "embeds": [{
      "title": "Deploy Status",
      "description": "Production deployment completed",
      "color": 3066993,
      "footer": {"text": "DevClaw Automation"}
    }]
  }'

# Colors: Red=15158332, Green=3066993, Blue=3447003, Yellow=16776960
```

## Tips

- Webhooks are easiest for simple notifications
- Bot API requires creating an app at https://discord.com/developers/applications
- Embed color is a decimal integer (not hex)
- Max message length: 2000 characters
- Rate limit: ~5 requests/second per channel

## Triggers

discord, send discord message, discord webhook, discord notification,
notify discord, discord bot
