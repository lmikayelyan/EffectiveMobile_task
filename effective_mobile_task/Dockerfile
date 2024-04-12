# Build Stage
FROM golang:1.22.2-alpine AS BuildStage
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8888
RUN go build -o /test main.go

FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /test /test
EXPOSE 8888
ENTRYPOINT ["/test"]

