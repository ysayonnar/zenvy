FROM --platform=$BUILDPLATFORM golang:1.26.0 AS builder

WORKDIR /app

COPY shared/go.mod shared/go.mod
COPY shared/go.sum shared/go.sum
COPY services/pipeline_service/go.mod services/pipeline_service/go.mod
COPY services/pipeline_service/go.sum services/pipeline_service/go.sum
COPY go.work ./
COPY go.work.sum ./
RUN go mod download -C shared && go mod download -C services/pipeline_service

COPY shared/ shared/
COPY services/pipeline_service/ services/pipeline_service/

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app ./services/pipeline_service/cmd/

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/services/pipeline_service/.env .

CMD [ "./app" ]
