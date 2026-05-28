FROM python:3.12 AS builder

WORKDIR /app

COPY requirements.txt .

RUN pip install --user -r requirements.txt

FROM python:3.12-slim

WORKDIR /app

COPY --from=builder /root/.local /root/.local

COPY . .

EXPOSE 8501

CMD ["python", "app.py"]

