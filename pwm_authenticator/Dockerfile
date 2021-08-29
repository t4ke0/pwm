FROM golang:1.16 as builder
WORKDIR /src/WORKDIR
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/serve/*


FROM alpine as final
WORKDIR /opt/bin
COPY --from=builder /src/WORKDIR/server .
ENTRYPOINT ["./server"]
