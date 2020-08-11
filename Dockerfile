FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git  make nodejs nodejs-npm gcc libc-dev

RUN mkdir -p /go/src/github.com/TaKeO90/pwm/

WORKDIR /go/src/github.com/TaKeO90/pwm/

COPY . /go/src/github.com/TaKeO90/pwm/
RUN make router
WORKDIR /go/src/github.com/TaKeO90/pwm/myfrontend/
RUN npm install && npm run build 

WORKDIR /go/src/github.com/TaKeO90/pwm/ssl-proxy
RUN make build

WORKDIR /go/src/github.com/TaKeO90/pwm/

VOLUME /go/src/github.com/TaKeO90/pwm/

EXPOSE 4430/tcp
EXPOSE 8080/tcp
EXPOSE 5000/tcp

CMD "./router";"./runFnt.sh";"./ssl-proxy/ssl-proxy"
