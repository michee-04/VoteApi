package controller

import (
	"encoding/json"
	"net/http"

	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
	"github.com/michee/micgram/pkg/utils/email"
)



func CreateUser(w http.ResponseWriter, r *http.Request){
	u := &model.User{}
  // w.Write([]byte("Please check your email to verify your account."))
	utils.ParseBody(r, u)
	user := u.CreateUser()
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	email.SendVerificationEmail(user)
	w.Write([]byte("Please check your email to verify your account."))

}

// hashedPassword, _ := utils.HashPassword(u.Password)
// 	token := hook.GenerateVerificationToken()
	
// 	user := User{
// 		Name: u.Name,
// 		Username: u.Username,
// 		Email: u.Email,
// 		Password: string(hashedPassword),
// 		EmailVerified: false,
// 		VerificationToken: token,
// 	}
// 	// email.SendVerificationEmail()