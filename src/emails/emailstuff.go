package emails

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(mnr string, name string) {
	m := gomail.NewMessage()

	// TODO
	// Set E-Mail sender
	m.SetHeader("From", "")

	// Set E-Mail receiver
	email := "s" + mnr + "@ba-sachens.de"
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Neue Prüfungsergebnisse verfügbar")

	// Set E-Mail body
	m.SetBody("text/html", "Hallo "+name+"\nSie haben neue Prüfungsergebnisse. Bitte prüfen Sie auf Campus Dual.")

	// TODO
	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "username", "password")

	// Because probably SSL/TLS won't be working
	// For production set to false
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Fatalln(err)
	}
	return
}
