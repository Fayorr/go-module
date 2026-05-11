# Stage 0: Build the application
FROM golang:1.26.2-alpine3.23 AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

FROM base AS builder
RUN CGO_ENABLED=0 go build -o /app/dockerize ./

FROM scratch
WORKDIR /app
COPY --from=builder /app/dockerize .
COPY --from=builder /app/banner ./banner
COPY --from=builder /app/templates ./templates


ENTRYPOINT [ "./dockerize" ]