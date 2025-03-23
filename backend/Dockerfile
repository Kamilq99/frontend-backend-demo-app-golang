FROM golang:1.24.1-bookworm

COPY . /app

WORKDIR /app

RUN go mod download

RUN go build -o serve

CMD ["./serve"]