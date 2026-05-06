# Stage 0: Build the application
FROM golang:1.26.2-alpine3.23 AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

FROM base AS builder
RUN go build -o /app/dockerize ./

FROM scratch
WORKDIR /app
COPY --from=builder /app/dockerize .
COPY --from=builder 
