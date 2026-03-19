# Makefile для проекта FastAPI + Worker + Redis + Postgres

# Переменные

DOCKER_COMPOSE = docker compose
SERVICE_FASTAPI = fastapi
SERVICE_WORKER = worker
SERVICE_REDIS = redis-1
SERVICE_POSTGRES = postgres

.PHONY: all up down logs clean restart shell

# По умолчанию — сборка и запуск
all: build up

# Собрать образы и понять их
up:
	$(DOCKER_COMPOSE) up --build

# Остановить все сервисы
down:
	$(DOCKER_COMPOSE) down

# Перезапустить сервисы (остановка + запуск)
restart: down up

# Посмотреть логи всех сервисов
logs:
	$(DOCKER_COMPOSE) logs -f

# Посмотреть логи конкретного сервиса
logs-fastapi:
	$(DOCKER_COMPOSE) logs -f $(SERVICE_FASTAPI)

logs-worker:
	$(DOCKER_COMPOSE) logs -f $(SERVICE_WORKER)

# Остановить и удалить образы, контейнеры и сети (очистка)
clean:
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans

# Открыть shell внутри контейнера fastapi
shell:
	$(DOCKER_COMPOSE) exec $(SERVICE_FASTAPI) /bin/bash




