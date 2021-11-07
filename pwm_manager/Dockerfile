FROM golang:1.17 as builder
WORKDIR /src/pwm_manager
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o service server/main.go


FROM alpine as final
WORKDIR /opt/bin
COPY --from=builder /src/pwm_manager/pw_generator/words.txt .
COPY --from=builder /src/pwm_manager/service .
ENTRYPOINT ["./service"]
