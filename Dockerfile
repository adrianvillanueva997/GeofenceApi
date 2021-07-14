# Multistage docker image building
# Build stage
FROM golang:1.16.6-alpine as build-env
RUN apk add --no-cache make
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

# Production stage
FROM alpine:3.14.0
WORKDIR /app
EXPOSE 3000
COPY --from=build-env /build/app .
RUN adduser -D appuser && mkdir logs && chown -R appuser:appuser /app && chmod 755 /app
USER appuser
ENTRYPOINT ./app
