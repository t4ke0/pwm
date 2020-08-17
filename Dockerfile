FROM golang:alpine AS backend
RUN apk update && apk upgrade && \
    apk add --no-cache bash git  make  gcc libc-dev ca-certificates
RUN mkdir -p /go/src/github.com/TaKeO90/pwm/
WORKDIR /go/src/github.com/TaKeO90/pwm/
COPY . /go/src/github.com/TaKeO90/pwm/
RUN make router
EXPOSE 8080/tcp
CMD ["./router"]


FROM node:12.18.3-alpine3.10 AS frontend
RUN mkdir myfrontend
WORKDIR myfrontend/
COPY ./myfrontend .
RUN npm install
EXPOSE 5000/tcp
CMD ["nmp start"]


FROM golang:alpine AS sslproxy
RUN apk update && apk upgrade && \
    apk add --no-cache bash make ca-certificates
RUN mkdir ssl-proxy/
WORKDIR ssl-proxy/
COPY ./ssl-proxy/ .
RUN ls -alth
RUN make build
VOLUME ssl-proxy/
EXPOSE 4430
CMD ["./ssl-proxy","-from","localhost:4430","-to","127.0.0.1:5000","-altnames","localhost"]
