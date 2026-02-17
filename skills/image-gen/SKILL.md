---
name: image-gen
description: "Generate images from text descriptions using DALL-E or GPT-image"
---
# Image Generation

The assistant will automatically use the generate_image tool when available.

## Setup

**API key** (DALL-E uses OpenAI; store in vault, never use `export`):
- `vault_save openai_api_key "sk-xxx"`
- Check: `vault_get openai_api_key` — key auto-injects as `OPENAI_API_KEY`

## Parameters
- prompt: Detailed description of the image
- size: 1024x1024 (default), 1024x1792 (portrait), 1792x1024 (landscape)
- quality: standard or hd

## Tips
- Be descriptive in prompts
- Specify style, mood, lighting for better results
