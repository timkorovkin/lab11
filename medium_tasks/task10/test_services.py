import pytest
import httpx
import subprocess
import time
import os

GO_URL = "http://localhost:8080"
PYTHON_URL = "http://localhost:8000"


@pytest.fixture(scope="session", autouse=True)
def start_services():
    env = os.environ.copy()
    env.update({
        "APP_ENV": "test",
        "APP_VERSION": "1.0.0",
        "GO_PORT": "8080",
        "PYTHON_PORT": "8000"
    })
    go_process = subprocess.Popen(
        ["go", "run", "main.go"],
        cwd=os.path.join(os.path.dirname(__file__), "go_service"),
        env=env
    )
    python_process = subprocess.Popen(
        ["uvicorn", "main:app", "--port", "8000"],
        cwd=os.path.join(os.path.dirname(__file__), "python_service"),
        env=env
    )
    time.sleep(5)
    yield
    go_process.terminate()
    python_process.terminate()


def test_go_health():
    response = httpx.get(f"{GO_URL}/health")
    assert response.status_code == 200


def test_python_health():
    response = httpx.get(f"{PYTHON_URL}/health")
    assert response.status_code == 200


def test_go_config():
    response = httpx.get(f"{GO_URL}/config")
    assert response.status_code == 200
    data = response.json()
    assert data["app_env"] == "test"
    assert data["app_version"] == "1.0.0"
    assert data["port"] == "8080"


def test_python_config():
    response = httpx.get(f"{PYTHON_URL}/config")
    assert response.status_code == 200
    data = response.json()
    assert data["app_env"] == "test"
    assert data["app_version"] == "1.0.0"
    assert data["port"] == "8000"