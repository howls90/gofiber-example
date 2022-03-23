FROM golang:1.17

WORKDIR /go/src/boilerplate

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

CMD ["air"]