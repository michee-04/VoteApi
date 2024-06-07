package hook

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/michee/micgram/pkg/model"
	"gopkg.in/gomail.v2"
)

func GenerateVerificationToken() string{
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func SendVerificationEmail(user *model.User){
	m := gomail.NewMessage()
	m.SetHeader("From", "nangmamichee0@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("../Verification/index.html", fmt.Sprintf("Click <a href=\"http://localhost:3000/verify?token=%s\">here</a> to verify your email address.", user.VerificationToken))

	d := gomail.NewDialer("smtp.gmail.com", 587, "nangmamichee0@gmail.com", "123456")


	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

