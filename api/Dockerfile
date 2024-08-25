FROM golang:1.21-alpine

WORKDIR /app/cmd/api

RUN go install github.com/cosmtrek/air@v1.28.0


COPY . .

RUN go mod download

CMD ["air"]