# Django Watch Brand — Minimal Site

Quick start:

1. Create a virtualenv and activate it

```bash
python -m venv .venv
source .venv/bin/activate   # or `.venv\Scripts\activate` on Windows
pip install -r requirements.txt
```

2. Run migrations and start server

```bash
python manage.py migrate
python manage.py runserver
```

3. Open http://127.0.0.1:8000/

This project contains a simple `catalog` app with a `Watch` model and basic templates. If no watches exist in the database, the site shows a small fallback sample list.
