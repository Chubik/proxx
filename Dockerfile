FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o proxx .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/proxx .
COPY assets/status.html assets/status.html
EXPOSE ${CONTAINER_PORT}
ENTRYPOINT ["./proxx"]
