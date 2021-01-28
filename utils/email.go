package utils

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(email, subject, content string) error {
	smtpHost := viper.Get("email.smtp_host")
	smtpPort := viper.Get("email.smtp_port")
	authEmail := viper.Get("email.auth_email")
	authPassword := viper.Get("email.auth_password")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", authEmail.(string))
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", content)

	dailer := gomail.NewDialer(
		smtpHost.(string),
		smtpPort.(int),
		authEmail.(string),
		authPassword.(string),
	)

	err := dailer.DialAndSend(mailer)

	return err
}