build-pipeline-service:
    @mkdir -p build
    @echo "Build pipeline service..."
    cd services/pipeline_service && go build -o ../../build/pipeline-service ./cmd
    @echo "Build completed"

clean:
    rm -rf build/*