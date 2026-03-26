build-pipeline-service:
    @mkdir -p build
    @echo "Build pipeline service..."
    cd pipeline_service && go build -o ../build/pipeline-service ./cmd
    @echo "Build completed"