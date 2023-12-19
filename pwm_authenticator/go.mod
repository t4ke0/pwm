module github.com/t4ke0/pwm/pwm_authenticator

go 1.16

require (
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.4 // indirect
	github.com/t4ke0/pwm/keys_manager v0.0.0-20210829112956-35b837e2af15
	github.com/t4ke0/pwm/pkg/common/http v0.0.0-20211121125600-19bddc4d5037
	github.com/t4ke0/pwm/pwm_db_api v0.0.0-20211114161425-6032255503d5
	golang.org/x/crypto v0.17.0
	google.golang.org/grpc v1.56.3
)
