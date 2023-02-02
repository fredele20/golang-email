package main

import (
	"fmt"
	packagegomail "golang-email/package-gomail"
)

func main() {
	fmt.Println("Testing email")
	// packagesmtp.SendEmail()
	packagegomail.SendEmail()
}
