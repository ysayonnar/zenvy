COMPOSE_FILE := "./infrastructure/docker-compose.yaml"

build-pipeline-service:
    @mkdir -p build
    @echo "Build pipeline service..."
    cd services/pipeline_service && go build -o ../../build/pipeline-service ./cmd
    @echo "Build completed"

pipeline_service_up_migrations:
    cd services/pipeline_service && goose up

up:
    docker compose -f {{ COMPOSE_FILE }} up --build -d

down:
    docker compose -f {{ COMPOSE_FILE }} down
