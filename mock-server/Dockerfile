FROM golang:1.17-alpine

WORKDIR /app

COPY ./app /app
COPY ./demo-fixtures /demo-fixtures

RUN go build -o fixtures-server

CMD ["./fixtures-server"]  