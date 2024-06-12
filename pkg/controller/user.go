package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/michee/micgram/pkg/email"
	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
)

var NewUser model.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	utils.ParseBody(r, u)
	user := u.CreateUser()
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	email.SendVerificationEmail(user)
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

	tmpl := template.Must(template.ParseFiles("../../templates/index.tmpl"))

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

	if creds.Password == "" {
		http.Error(w, "Password cannot be empty", http.StatusBadRequest)
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

	// Debug log pour vérifier les informations de l'utilisateur avant de vérifier le mot de passe
	fmt.Printf("User found: %v, Password hash: %s\n", user.Email, user.Password)
	fmt.Printf("Password provided for login: %s\n", creds.Password)

	// Comparer le mot de passe fourni avec le mot de passe haché
	if !utils.CheckPasswordHash(creds.Password, user.Password) {
		// Ajoutez un journal pour indiquer que la vérification du mot de passe a échoué
		fmt.Println("Password does not match")
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Ajoutez un journal pour indiquer que la vérification du mot de passe a réussi
	fmt.Println("Password match successful")

	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write([]byte(`{"message": "Login successful"}`))
}




func GetUser(w http.ResponseWriter, r *http.Request) {
	NewUser := model.GetAllUser()
	res, _ := json.Marshal(NewUser)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	user, _ := model.GetUserById(userId)


	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



// reinitilisatiom du mot de passe par email echoue
/**
	func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		var user model.User
		if err := model.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err := user.GenerateResetToken(); err != nil {
			http.Error(w, "Failed to generate reset token", http.StatusInternalServerError)
			return
		}


		email.SendResetPasswordEmail(&user)

		tmpl := template.Must(template.ParseFiles("../../templates/forgotPass.tmpl"))

		w.Header().Set("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)
	}



	func ResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
		// Parse le fichier HTML
		tmpl := template.Must(template.ParseFiles("../../templates/forgotPass.tmpl"))

		// Définir le content-type à text/html
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		// Exécuter le template avec une structure de données vide
		tmpl.Execute(w, nil)
	}


	func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Token       string `json:"token"`
			NewPassword string `json:"newPassword"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		user, err := model.FindUserByResetToken(req.Token)
		if err != nil {
			http.Error(w, "Invalid or expired reset token", http.StatusUnauthorized)
			return
		}

		if err := user.UpdatePassword(req.NewPassword); err != nil {
			http.Error(w, "Failed to update password", http.StatusInternalServerError)
			return
		}

		res, _ := json.Marshal(user)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		w.Write([]byte(`{"message": "Password updated successfully"}`))
	}
**/

func UpdateUser(w http.ResponseWriter, r *http.Request){
	userUpdate := &model.User{}
	utils.ParseBody(r, &userUpdate)
	userId := chi.URLParam(r, "userId")
	userDetails, db := model.GetUserById(userId)

	if userDetails == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if userUpdate.Email != "" {
		userDetails.Email = userUpdate.Email
		email.SendVerificationEmail(userDetails)
	}

	if userUpdate.Password != "" {
		hashedPassword, _ := utils.HashPassword(userUpdate.Password)
		userDetails.Password = hashedPassword
	}

	if userUpdate.Username != ""{
		userDetails.Username = userUpdate.Username
	}

	if userUpdate.Name != ""{
		userDetails.Name = userUpdate.Name
	}

	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
  w.Header().Set("content-type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}


func DeleteUser(w http.ResponseWriter, r *http.Request){
	userId := chi.URLParam(r, "userId")
	user := model.DeleteUserId(userId)
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}