---
name: pip-poetry
description: "Python package management with pip, poetry, and pipenv"
---
# pip / Poetry / pipenv

Use the **bash** tool for Python dependency management.

## Setup

1. **Check if installed:**
   ```bash
   command -v python3 && python3 --version
   command -v pip && pip --version
   command -v poetry && poetry --version
   ```

2. **Install:**
   ```bash
   # Python & pip
   # macOS
   brew install python3

   # Ubuntu / Debian
   sudo apt update && sudo apt install -y python3 python3-pip

   # Poetry (official installer)
   curl -sSL https://install.python-poetry.org | python3 -
   ```

## pip
```bash
pip install <pkg>
pip install -r requirements.txt
pip freeze > requirements.txt
pip list --outdated
pip audit
```

## Poetry
```bash
poetry init
poetry add <pkg>
poetry add --group dev <pkg>
poetry install
poetry lock
poetry show --tree
poetry run python script.py
poetry build
poetry publish
```

## pipenv
```bash
pipenv install <pkg>
pipenv install --dev <pkg>
pipenv lock
pipenv run python script.py
pipenv check
```

## Virtual Environments
```bash
python -m venv .venv
source .venv/bin/activate
deactivate
```

## Tips
- Detect tool by file: requirements.txt (pip), pyproject.toml (poetry), Pipfile (pipenv)
- Use pip audit or safety check for vulnerability scanning
- Use poetry for new projects (better dependency resolution)
