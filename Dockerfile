FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git curl

RUN mkdir -p $GOPATH/src/github.com/tarektouati/talking-assistant
WORKDIR $GOPATH/src/github.com/tarektouati/talking-assistant
COPY . ./
RUN go build -o talking-assistant cmd/talking-assistant/main.go

FROM alpine
WORKDIR /root
COPY --from=builder /go/src/github.com/tarektouati/talking-assistant/talking-assistant .
ENTRYPOINT ["./talking-assistant"]