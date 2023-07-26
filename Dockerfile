FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./...

RUN cd ./cli && CGO_ENABLED=0 GOOS=linux go build -o 15

ENTRYPOINT ["./cli/15"]
