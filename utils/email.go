package utils

import (
	"blogApp/config"
	"fmt"
	"net/smtp"
)

func SendApprovalEmail(to, name string) error {
	auth := smtp.PlainAuth("", config.USER_EMAIL, config.USER_PASS, "smtp.gmail.com")

	msg := []byte(fmt.Sprintf("Subject: Registration Successful\r\n\r\nDear %s,\n\nYour Registration is Successful. You can now log in.\n\nRegards,\nBLOG App", name))

	return smtp.SendMail("smtp.gmail.com:587", auth, config.USER_EMAIL, []string{to}, msg)
}
