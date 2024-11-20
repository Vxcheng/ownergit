package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/wneessen/go-mail"
	"gopkg.in/gomail.v2"
)

func main() {
	log.Println("Starting...")
	standardMail()
}

const (
	host               = "smtp.qq.com"
	port               = 587
	username, password = "834459936@qq.com", "rhosqnxwalcubaig"
)

func goMail() {

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", username)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <br/><b>world?</b>")

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Could not send email: %v", err)
		return
	}
	log.Println("Email sent!")
}

func goEmail() {

	e := email.NewEmail()
	e.From = username
	e.To = []string{username}
	e.Subject = "Awesome web"
	e.Text = []byte("Text Body is, of course, supported!")
	err := e.Send(fmt.Sprintf("%s:%d", host, port), smtp.PlainAuth("", username, password, host))
	if err != nil {

		log.Fatal("Send: %v ", err)
		return
	}
	log.Println("Email sent!")

}

func Mail() {
	m := mail.NewMsg()
	if err := m.From(username); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(username); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("This is my first mail with go-mail!")
	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")
	c, err := mail.NewClient(host, mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(username), mail.WithPassword(password))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}

	log.Println("Email sent!")

}

func standardMail() {
	auth := smtp.PlainAuth("", username, password, host)
	value := fmt.Sprintf(`"To: %s\r\n" +
	"Subject: Hello!\r\n" +
	"\r\n" +
	"This is the email body.\r\n"`, username)
	msg := []byte(value)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port),
		auth, username,
		[]string{username},
		msg)
	if err != nil {
		log.Println("failed to send email:", err)
		return
	}
	log.Println("Email sent!")
}
