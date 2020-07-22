FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git gcc libc-dev
RUN mkdir -p ~/go/src/
WORKDIR ~/go/src/
RUN cd ~/go/src/ && git clone https://github.com/TaKeO90/pwm.git && cd pwm && go get -d
RUN go build  -o launcher ~/go/src/pwm/launcher.go
CMD ["./launcher"]
EXPOSE 8080/tcp
