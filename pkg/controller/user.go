package controller

import (
	"encoding/json"
	"net/http"

	"github.com/michee/micgram/pkg/hook"
	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
)



func CreateUser(w http.ResponseWriter, r *http.Request){
	u := &model.User{}
	hashedPassword, _ := utils.HashPassword(u.Password)
	token := hook.GenerateVerificationToken()
	
	user := model.User{
		Name: u.Name,
		Username: u.Username,
		Email: u.Email,
		Password: string(hashedPassword),
		EmailVerified: false,
		VerificationToken: token,
	}

	// hook.SendVerificationEmail(&user)

  w.Write([]byte("Please check your email to verify your account."))
	utils.ParseBody(r, user)
	users := user.CreateUser()
	res, _ := json.Marshal(users)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}