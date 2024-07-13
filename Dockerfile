FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

RUN golangci-lint run --issues-exit-code=0

RUN go build -o main ./cmd/golang-exercise/main.go

EXPOSE 8080

CMD ["./main"]
