FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY byte-battle_backend/go.mod byte-battle_backend/go.sum ./
RUN go mod download

COPY ./byte-battle_backend .
RUN go build -v cmd/apiserver/main.go

CMD ["./main"]

