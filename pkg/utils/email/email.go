package email

import (
	"fmt"

	"github.com/michee/micgram/pkg/model"
	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(user *model.User) {
	m := gomail.NewMessage()
	m.SetHeader("From", "voteprojet@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Veuillez activer votre compte")
	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://localhost:3000/verify?token=%s\">here</a> to verify your email address.", user.VerificationToken))

	d := gomail.NewDialer("smtp.gmail.com", 587, "voteprojet@gmail.com", "jmbd aicq hdov mvyq")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
