package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
	"github.com/michee/micgram/pkg/utils/email"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	utils.ParseBody(r, u)
	user := u.CreateUser()
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	email.SendVerificationEmail(user)
	w.Write([]byte("Please check your email to verify your account."))
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	user, err := model.FindUserByToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = user.Verify()
	if err != nil {
		http.Error(w, "Unable to verify user", http.StatusInternalServerError)
		return
	}

	// Utilisez un chemin absolu pour charger le fichier template
	absPath, err := filepath.Abs("templates/index.tmpl")

	if err != nil {
		log.Printf("Error getting absolute path: %v", err)
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(absPath)
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user model.User
	if err := model.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !user.CanLogin() {
		http.Error(w, "Email not verified", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Login successful, handle session creation, etc.
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write([]byte(`{"message": "Login successful"}`))
}
