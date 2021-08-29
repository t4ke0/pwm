FROM golang:1.16 as builder
WORKDIR /src/WORKDIR
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./cmd/main.go

FROM alpine as final
WORKDIR /opt/bin
COPY --from=builder /src/WORKDIR/schema.sql .
COPY --from=builder /src/WORKDIR/service .
ENV SCHEMA_FILE_PATH=/opt/bin/schema.sql
RUN ls -alth
ENTRYPOINT ["./service"]
