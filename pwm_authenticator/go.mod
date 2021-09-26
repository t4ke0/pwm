module github.com/t4ke0/pwm/pwm_authenticator

go 1.16

require (
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/google/uuid v1.1.2
	github.com/lib/pq v1.10.3 // indirect
	github.com/t4ke0/pwm/keys_manager v0.0.0-20210829112956-35b837e2af15
	github.com/t4ke0/pwm/pwm_db_api v0.0.0-20210926111908-a8c122c0aebd
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.39.0
)
