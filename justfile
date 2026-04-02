COMPOSE_FILE := "./infrastructure/docker-compose.yaml"

pipeline_service_up_migrations:
    cd services/pipeline_service && goose up

up:
    docker compose -f {{ COMPOSE_FILE }} -p zenvy up --build -d

down:
    docker compose -f {{ COMPOSE_FILE }} -p zenvy down
