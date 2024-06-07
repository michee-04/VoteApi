package email

import (
	// "fmt"
	"log"
	"net/smtp"

	"github.com/michee/micgram/pkg/model"
	// "gopkg.in/gomail.v2"
)


func SendVerificationEmail(user *model.User){

	username := "1007649af53941"
	password := "4f620ba9b62604"
	smtpHost := "sandbox.smtp.mailtrap.io"

	auth := smtp.PlainAuth("", username, password, smtpHost)


	from := user.Email
	to := []string{"mailtrap.foo@gmail.com"}
	message := []byte("To: mailtrap.foo@gmail.com\r\n"+
  "From: " + user.Email + ",\r\n"+  // Add comma here
  "Subject: Click <a href=\"http://localhost:3000/verify?token=%s\">here</a> to verify your email address")

	// "Subject: Click <a href=\"http://localhost:3000/verify?token=%s\">here</a> to verify your email address", user.VerificationToken


	smtpUrl := smtpHost + ":587"

	err := smtp.SendMail(smtpUrl, auth, from, to, message)

	if err != nil {
		log.Fatal(err)
	}


	// m := gomail.NewMessage()
	// m.SetHeader("From", "mailtrap.foo@gmail.com")
	// m.SetHeader("To", "nangmamichee2@gmail.com")
	// m.SetAddressHeader("Cc", user.Email, user.Username)
	// m.SetHeader("Subject", "Email Verification")
	// m.SetBody("../Verification/index.html", fmt.Sprintf("Click <a href=\"http://localhost:3000/verify?token=%s\">here</a> to verify your email address.", user.VerificationToken))

	// d := gomail.NewDialer("smtp.gmail.com", 587, "mailtrap.foo@gmail.com", "4f620ba9b62604")


	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }
}
