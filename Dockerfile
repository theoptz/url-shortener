FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
RUN go build -o cmd/server/server cmd/server/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/cmd/server/server ./

EXPOSE 3000
CMD ./server --host 0.0.0.0 --port 3000
