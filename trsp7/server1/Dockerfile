FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/

EXPOSE 10001
EXPOSE 3000

CMD ["./main"]