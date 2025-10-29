FROM golang:1.25.1-alpine AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN mkdir -p /app/bin

RUN go build -o /app/bin/api ./cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN adduser -D -s /bin/sh appuser

COPY --from=builder /app/bin/api /bin/api
COPY --from=builder /app/web /web

USER appuser

EXPOSE 9999
CMD ["/bin/api"]