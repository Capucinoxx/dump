FROM golang:1.22.1-alpine

COPY src/ /app

WORKDIR /app
RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]

