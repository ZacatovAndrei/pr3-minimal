FROM golang:1.19.3-alpine

WORKDIR /server

COPY *.* ./

EXPOSE 8080


RUN go build -o server

EXPOSE 8080

ENTRYPOINT  ["/server/server"]
