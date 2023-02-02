package main

import (
	"fmt"
	htmltemplate "golang-email/html-template"
)

func main() {
	fmt.Println("Testing email")
	// packagesmtp.SendEmail()
	// packagegomail.SendEmail()
	htmltemplate.SendEmail()
}
