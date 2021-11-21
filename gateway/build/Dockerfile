FROM golang:1.17 as builder
WORKDIR /src
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o server /src/cmd/serve/main.go


FROM alpine as final
WORKDIR /opt/bin
COPY --from=builder /src/server .
ENTRYPOINT ["./server"]
