package gmailx

import (
	"cleaningservice/common/variables"
	"log"

	"gopkg.in/gomail.v2"
)

const emailEnd string = "<p>If you have any questions and concerns, you can kindly reply to this email</p><br><p>Kind regards,</p><br>"
const emailSign string = variables.Business_name + " Support Team"

func SendToClient(target string, subject string, content string, attached string) {
	// Send email
	m := gomail.NewMessage()
	m.SetHeader("From", variables.QME_email)
	m.SetHeader("To", target)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content+emailEnd+emailSign)
	m.Attach(attached)

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, variables.QME_email, variables.QME_password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}

func SendToReception(subject string, content string) {
	// Send email
	m := gomail.NewMessage()
	m.SetHeader("From", variables.QME_email)
	m.SetHeader("To", variables.Reception_email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content+emailEnd+emailSign)

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, variables.QME_email, variables.QME_password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
