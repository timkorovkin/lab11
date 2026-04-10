## Task 2 — Написать Dockerfile для Go-приложения с многоэтапной сборкой

Инструмент: Claude (claude.ai)
Промпт: "Написать Dockerfile для Go-приложения с многоэтапной сборкой"

Что сделано:
1. Создан `task2/handlers/handlers.go` — 2 эндпоинта: `GET /status`, `GET /items`. Хранит список items в памяти.
2. Создан `task2/main.go` — Gin-сервер на порту 8080.
3. Создан `task2/Dockerfile` — многоэтапная сборка: этап 1 (`golang:1.23-alpine`) компилирует бинарник, этап 2 (`scratch`) содержит только бинарник. Итоговый размер образа: 17.3MB вместо ~500MB.
4. Создан `task2/main_test.go` — 2 Go-теста через `net/http/httptest`, проверяют статус ответов эндпоинтов.

Результат: все работает, 2/2 Go-тестов прошли. Образ собран и запущен, эндпоинты отвечают корректно.

Итого:
* Количество промптов: 1
* Что пришлось исправлять вручную: зафиксировать версию Gin v1.9.1 в go.mod — последняя версия требовала Go 1.25, которого нет в alpine-образе; исправить версию Go в go.mod с 1.21 на 1.23
* Время: ~50 мин

## Task 8 — Добавить healthcheck для каждого сервиса

Инструмент: Claude (claude.ai)
Промпт: "Добавить healthcheck для Go и Python сервисов в Docker"

Что сделано:
1. Создан `task8/go_service/handlers/handlers.go` — 2 эндпоинта: `GET /health`, `GET /items`.
2. Создан `task8/go_service/main.go` — Gin-сервер на порту 8080.
3. Создан `task8/go_service/Dockerfile` — многоэтапная сборка на базе `golang:1.23-alpine` и `alpine:3.19`. Установлен curl для healthcheck. Инструкция `HEALTHCHECK` проверяет `/health` каждые 30 секунд.
4. Создан `task8/python_service/main.py` — FastAPI-сервис на порту 8000 с эндпоинтами `GET /health` и `GET /items`.
5. Создан `task8/python_service/Dockerfile` — образ на базе `python:3.11-slim`. Инструкция `HEALTHCHECK` проверяет `/health` через httpx каждые 30 секунд.
6. Создан `task8/docker-compose.yml` — запускает оба сервиса, healthcheck прописан для каждого, `python_service` запускается только после того как `go_service` становится `healthy`.
7. Создан `task8/main_test.go` — 2 Go-теста через `net/http/httptest`.
8. Создан `task8/test_services.py` — 4 pytest-теста, проверяют healthcheck и эндпоинты обоих сервисов.

Результат: все работает, 2/2 Go-тестов и 4/4 pytest-тестов прошли. Оба контейнера показывают статус `healthy` в `docker ps`.

Итого:
* Количество промптов: 1
* Что пришлось исправлять вручную: зафиксировать версию Gin v1.9.1 в go.mod; увеличить time.sleep с 3 до 5 секунд в тестах; переименовать docker_compose.yml в docker-compose.yml
* Время: ~50 мин

## Task 10 — Использовать переменные окружения для конфигурации

Инструмент: Claude (claude.ai)
Промпт: "Использовать переменные окружения для конфигурации Go и Python сервисов в Docker"

Что сделано:
1. Создан `task10/go_service/handlers/handlers.go` — 3 эндпоинта: `GET /health`, `GET /items`, `GET /config`. Эндпоинт `/config` возвращает текущие значения переменных окружения `APP_ENV`, `APP_VERSION`, `GO_PORT`.
2. Создан `task10/go_service/main.go` — Gin-сервер, порт берётся из переменной окружения `GO_PORT`, fallback на `8080`.
3. Создан `task10/go_service/Dockerfile` — многоэтапная сборка, HEALTHCHECK использует `${GO_PORT:-8080}`.
4. Создан `task10/python_service/main.py` — FastAPI-сервис с эндпоинтом `/config`, читает `APP_ENV`, `APP_VERSION`, `PYTHON_PORT` через `os.getenv()` с fallback-значениями.
5. Создан `task10/python_service/Dockerfile` — порт берётся из `${PYTHON_PORT:-8000}`.
6. Создан `task10/.env` — содержит `APP_ENV`, `APP_VERSION`, `GO_PORT`, `PYTHON_PORT`.
7. Создан `task10/docker-compose.yml` — оба сервиса читают переменные из `.env` через `env_file`, порты задаются через переменные окружения.
8. Создан `task10/main_test.go` — 3 Go-теста, включая проверку эндпоинта `/config` с заданными переменными окружения через `os.Setenv`.
9. Создан `task10/test_services.py` — 4 pytest-теста, проверяют что оба сервиса правильно читают переменные окружения.

Результат: все работает, 3/3 Go-тестов и 4/4 pytest-тестов прошли. Оба сервиса корректно читают конфигурацию из переменных окружения.

Итого:
* Количество промптов: 1
* Что пришлось исправлять вручную: зафиксировать версию Gin v1.9.1 в go.mod
* Время: ~40 мин