package htmltemplate

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/gomail.v2"
)


func SendEmail() {
	file, err := ioutil.ReadFile("./html-template/mail.html")
	if err != nil {
		fmt.Printf("Could not open the file due to %v", err.Error())
		return
	}

	//
	msg := gomail.NewMessage()
	msg.SetHeader("From", "victorisholaoladele@gmail.com")
	msg.SetHeader("To", "frederickvictor93@gmail.com")
	msg.SetHeader("Subject", "Testing...")
	msg.SetBody("text/html", string(file))

	dialer := gomail.NewDialer("smtp.gmail.com",587,"victorisholaoladele@gmail.com", "")

	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Printf("Could not send the Email due to %v", err.Error())
		return
	}

	fmt.Println("Email sent successfully!")
}