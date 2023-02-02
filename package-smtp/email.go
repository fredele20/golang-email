package packagesmtp

import (
	"fmt"
	"net/smtp"
)

func SendEmail() {
	// sender data.
	from := "victorisholaoladele@gmail.com"
	password := ""

	// Receiver email address
	to := []string{
		"frederickvictor93@gmail.com",
	}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	subject := "Testing...\n"
	body := "This is a test email message"
	message := []byte(subject + body)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// sending email

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Email sent successfully")
}
