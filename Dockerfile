FROM golang:1.19.6 as base

VOLUME [ "/opt/app" ]

WORKDIR /opt/app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]
