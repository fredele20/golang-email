package packagegomail

import "gopkg.in/gomail.v2"


func SendEmail() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "victorisholaoladele@gmail.com")
	msg.SetHeader("To", "frederickvictor93@gmail.com")
	msg.SetHeader("Subject", "Testing...")
	msg.SetBody("text/html", "<b>This is the body of the test mail</b>")

	n := gomail.NewDialer("smtp.gmail.com",587,"victorisholaoladele@gmail.com", "")

	// send mail 
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}