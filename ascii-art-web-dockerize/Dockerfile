# Stage 0: Build the application
FROM golang:1.26.2-alpine3.23 AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

FROM base AS builder
RUN CGO_ENABLED=0 go build -o /app/dockerize ./

FROM scratch

LABEL maintainers="fayokunmiosho@gmail.com, emmanuelsensai@gmail.com"
LABEL project="Ascii Art Dockerize"
LABEL description="Web server that converts text to Ascii-Art"

WORKDIR /app
COPY --from=builder /app/dockerize ./
COPY --from=builder /app/banner ./banner
COPY --from=builder /app/templates ./templates

EXPOSE 4000

ENTRYPOINT [ "./dockerize" ]