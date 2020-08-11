
[![Go Report Card](https://goreportcard.com/badge/github.com/TaKeO90/pwm)](https://goreportcard.com/report/github.com/TaKeO90/pwm)

# SO FAR || ROADMAP
```
[+] Register & login 
[+] generating encryption key for the user for encrypting and decrypting credentials 
[+] Save & Show User Credentials 
[+] Generating an encryption key for the server when starting the launcher for the first time , in case the launcher have found the server key it skips this phase 
[+] Decypting user's key in case he need ot show , add or update his credentials
[+] Support REST API with JSON
[+] Get better score in go report card
[] make a docker image
[] should support kubernetes

```



# Usage

```sh
git clone --recurse-submodules  https://github.com/TaKeO90/pwm.git
make -B
./launcher

```
OR run it on Docker

```sh
docker build -t pwm:v1 .
docker run -p 4430:4430 -v <host direct>:/go/src/github.com/TaKeO90/pwm/ --name pwm -d pwm:v1

```
