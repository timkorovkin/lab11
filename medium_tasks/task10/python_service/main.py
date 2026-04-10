import os
from fastapi import FastAPI

app = FastAPI()


@app.get("/health")
async def health():
    return {"status": "healthy"}


@app.get("/items")
async def get_items():
    return [
        {"id": 1, "name": "Item One"},
        {"id": 2, "name": "Item Two"},
        {"id": 3, "name": "Item Three"},
    ]


@app.get("/config")
async def get_config():
    return {
        "app_env": os.getenv("APP_ENV", "development"),
        "app_version": os.getenv("APP_VERSION", "1.0.0"),
        "port": os.getenv("PYTHON_PORT", "8000"),
    }