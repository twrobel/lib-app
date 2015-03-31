package main

import (
        "log"
        "net/smtp"
)

func main() {
	to := "tenrally@gmail.com"
	from := "lib.reminder.app@gmail.com"
	pwd := "haselko123"
	subject := "test"
	msg := "test email"
	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + msg
   	auth := smtp.PlainAuth("", from, pwd, "smtp.gmail.com")
   	err := smtp.SendMail("smtp.gmail.com:587", auth, from,
   		[]string{to},[]byte(body))
   	if err != nil {
    	log.Fatal(err)
   	}
}