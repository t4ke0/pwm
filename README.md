
[![Go Report Card](https://goreportcard.com/badge/github.com/TaKeO90/pwm)](https://goreportcard.com/report/github.com/TaKeO90/pwm)

# SO FAR || ROADMAP
```
- [x] Register & login 
- [x] generating encryption key for the user for encrypting and decrypting credentials 
- [x] Save & Show User Credentials 
- [x] Generating an encryption key for the server when starting the launcher for the first time , in case the launcher have found the server key it skips this phase 
- [x] Decypting user's key in case he need ot show , add or update his credentials
- [x] Support REST API with JSON
- [x] Get better score in go report card
- [x] make a docker image for each service
- [ ] deploy those images using kubernetes

```



# Usage

```sh
git clone --recurse-submodules  https://github.com/TaKeO90/pwm.git
make -B
./launcher

```
OR run it on Docker

```sh
docker build --target backdend -t pwmbackend:v1 .
docker build --target frontend -t pwmfrontend:v1 .
docker build --target sslproxy -t sslproxy:v1 .

TODO finish this by implementing kubernetes deployments.....
```
