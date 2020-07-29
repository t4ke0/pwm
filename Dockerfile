FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git gcc libc-dev make

RUN mkdir -p /go/src/github.com/TaKeO90/pwm/

WORKDIR /go/src/github.com/TaKeO90/

RUN git clone https://github.com/TaKeO90/pwm.git

COPY . /go/src/github.com/TaKeO90/ 

#RUN go build -o launcher launcher.go
RUN make build

CMD ["./launcher"]

EXPOSE 4430
