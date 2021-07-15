FROM golang:1.13.4-alpine3.10 As builder
RUN apk update && apk add --no-cache git

WORKDIR /go/src/rest-api/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go get  github.com/swaggo/swag/cmd/swag \
    && go get  github.com/swaggo/gin-swagger \
    && go get  github.com/swaggo/files \
    && go get  github.com/alecthomas/template \
    && go get  github.com/swaggo/swag \
    && go get  github.com/pilu/fresh


ENTRYPOINT swag init && fresh