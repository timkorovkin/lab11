import pytest
import httpx
import subprocess
import time
import os

GO_URL = "http://localhost:8080"
PYTHON_URL = "http://localhost:8000"


@pytest.fixture(scope="session", autouse=True)
def start_services():
    go_process = subprocess.Popen(
        ["go", "run", "main.go"],
        cwd=os.path.join(os.path.dirname(__file__), "go_service")
    )
    python_process = subprocess.Popen(
        ["uvicorn", "main:app", "--port", "8000"],
        cwd=os.path.join(os.path.dirname(__file__), "python_service")
    )
    time.sleep(5)
    yield
    go_process.terminate()
    python_process.terminate()


def test_go_health():
    response = httpx.get(f"{GO_URL}/health")
    assert response.status_code == 200
    assert response.json()["status"] == "healthy"


def test_python_health():
    response = httpx.get(f"{PYTHON_URL}/health")
    assert response.status_code == 200
    assert response.json()["status"] == "healthy"


def test_go_items():
    response = httpx.get(f"{GO_URL}/items")
    assert response.status_code == 200
    assert len(response.json()) == 3


def test_python_items():
    response = httpx.get(f"{PYTHON_URL}/items")
    assert response.status_code == 200
    assert len(response.json()) == 3