FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/

EXPOSE 10002
EXPOSE 10003

CMD ["./main"]