FROM golang:1.16 as builder
WORKDIR /go/src/github.com/t4ke0/pwm/keys_manager
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o service /go/src/github.com/t4ke0/pwm/keys_manager/server/server.go

FROM alpine as final
WORKDIR /opt/bin
COPY --from=builder /go/src/github.com/t4ke0/pwm/keys_manager/service .
COPY --from=builder /go/src/github.com/t4ke0/pwm/keys_manager/common/words.txt .
ENTRYPOINT ["./service"]
