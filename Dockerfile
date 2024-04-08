FROM golang:1.21.6-alpine

WORKDIR /app

COPY . . 

RUN go get
RUN go build -o main .

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
## CMD ["go", "run", "main.go"]