FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/api


FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/user-service .
USER nonroot:nonroot

ENTRYPOINT ["/app/user-service"]
