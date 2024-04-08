FROM golang:1.21.6

WORKDIR /app

COPY . . 

RUN go get
RUN go build -o main .

ENTRYPOINT [ "/app/main" ]