package emailsender

import (
	"net/smtp"
	"os"
)

//EmailInfo struct holds email info
type EmailInfo struct {
	from       string
	msg        []byte
	recipients []string
}

var (
	password  string = os.Getenv("MAILPW")
	emailaddr string = os.Getenv("MAILADDR")
	hostname  string = "smtp.gmail.com"
	port      string = ":587"
)

//EmailAuth email authentication
func EmailAuth() smtp.Auth {
	auth := smtp.PlainAuth("", emailaddr, password, hostname)
	return auth
}

//SendCode sending code that we generate to the user
func SendCode(msg, email string) (bool, error) {
	info := &EmailInfo{}
	info.from, info.msg, info.recipients = "pwm.noreply", []byte(msg), []string{email}
	auth := EmailAuth()
	err := smtp.SendMail(hostname+port, auth, info.from, info.recipients, info.msg)
	if err != nil {
		return false, err
	}
	return true, nil
}
