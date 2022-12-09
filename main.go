package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

func main() {

	sender := "yourmail@example.com"

	to := []string{
		"receipt@example.com",
		"yourmail@example.com",
		"cc1@example.com",
		"cc2@example.com",
	}

	cc := []string{
		"cc1@example.com",
		"cc2@example.com",
	}

	bcc := []string{
		"yourmail@example.com",
	}

	user := "yourmail@example.com"
	password := "yourpassword"

	subject := "作業開始(yourname)"
	body := `<div>おはようございます。</div><div>作業を開始します。 </div>`

	request := Mail{
		Sender:  sender,
		To:      to,
		Cc:      cc,
		Bcc:     bcc,
		Subject: subject,
		Body:    body,
	}

	addr := "mail.airfolc.co.jp:587"
	host := "mail.airfolc.co.jp"

	msg := BuildMessage(request)
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}

func BuildMessage(mail Mail) string {
	msg := "Message-Id: <your message id>\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
	}

	if len(mail.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
