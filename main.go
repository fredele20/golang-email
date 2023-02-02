package main

import (
	"fmt"
	packagesmtp "golang-email/package-smtp"
)


func main() {
	fmt.Println("Testing email")
	packagesmtp.SendEmail()
}