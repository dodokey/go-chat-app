FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.28.0

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go mod vendor
WORKDIR /app/cmd/api

CMD ["air"]