FROM golang:1.25.5-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \ 
    -ldflags='-w -s -extldflags "-static"' \
    -o /app/main \
    ./cmd/api/

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

# Create non-root user for security (i guess)
RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser

WORKDIR /home/appuser

COPY --from=builder --chown=appuser:appgroup /app/main .

# SWITCH TO NON-ROOT USER
USER appuser

EXPOSE 8080

CMD ["./main"]