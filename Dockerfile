FROM golang:1.21.0-alpine3.18

EXPOSE 8080 

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

CMD ["./main"]