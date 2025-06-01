package smtp

import (
	"fmt"
	"net/smtp"
	"os"
)

func Send(to, subject, body string) error {
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailUser := os.Getenv("EMAIL_USER")
	emailPass := os.Getenv("EMAIL_PASS")

	auth := smtp.PlainAuth("", emailUser, emailPass, emailHost)

	msg := []byte("Subject: " + subject + "\r\n" +
		"\r\n" + body)

	err := smtp.SendMail(
		emailHost+":"+emailPort,
		auth,
		emailUser,
		[]string{to},
		msg,
	)
	if err != nil {
		return fmt.Errorf("SMTP error: %w", err)
	}
	return nil
}
