# Multistage docker image building
# Build stage
FROM golang:1.16.3-alpine as build-env
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app .

# Production stage
FROM alpine:3.13.4
WORKDIR /app
EXPOSE 3000
COPY --from=build-env /build/app .
RUN adduser -D appuser && mkdir logs && chown -R appuser:appuser /app && chmod 755 /app
USER appuser
ENTRYPOINT ./app
