FROM --platform=$BUILDPLATFORM golang:1.26.0 AS builder

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./ TODO: remove as need dependencies
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app ./cmd/

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
#COPY --from=builder /app/.env .

CMD [ "./app" ]