package gomailx

import (
	"cleaningservice/common/constantx"
	"log"

	"gopkg.in/gomail.v2"
)

const emailEnd string = "<p>If you have any questions and concerns, you can kindly reply to this email</p><br><p>Kind regards,</p><br>"
const emailSign string = constantx.Business_name + " Support Team"

func Send(target string, subject string, content string, attached string) error {
	// Send email
	m := gomail.NewMessage()
	m.SetHeader("From", constantx.QME_email)
	m.SetHeader("To", target)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content+emailEnd+emailSign)

	if len(attached) > 0 {
		m.Attach(attached)
	}

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, constantx.QME_email, constantx.QME_password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
