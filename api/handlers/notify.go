package handlers

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"io"
	"log"

	gomail "gopkg.in/mail.v2"
)

// Person is model
type Person struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	To       string `json:"to,omitempty"`
	Message  string `json:"message,omitempty"`
}

func notifyByEmail(templateName string, p *Person) {

	m := gomail.NewMessage()

	m.SetHeader("From", p.Email)
	m.SetHeader("To", p.To)

	m.SetHeader("Subject", "Gophers GO!")

	t := template.Must(template.ParseFiles(fmt.Sprintf("templates/%s", templateName)))
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, p)
	})

	log.Println(p.Password + p.Email)
	d := gomail.NewDialer("smtp.gmail.com", 587, p.Email, p.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
