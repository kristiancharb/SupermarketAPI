FROM golang:1.16.2-alpine3.13

RUN mkdir /app
RUN mkdir /app/src
RUN mkdir /app/bin

ADD . /app/src
WORKDIR /app/src

EXPOSE 8080

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o /app/bin/main .

CMD ["/app/bin/main"]