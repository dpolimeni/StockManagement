FROM golang:1.21.6-alpine

WORKDIR /app

COPY . . 

RUN go get
RUN go build -o main .

ENTRYPOINT [ "/app/main" ]