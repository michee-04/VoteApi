package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/michee/micgram/pkg/controller"
)

const port = ":3000"

func main (){
	r := chi.NewRouter()
	
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	// Authentification
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", controller.CreateUser)
		r.Get("/verify", controller.VerifyHandler)
		r.Post("/login", controller.LoginHandler)
		// r.Post("/forgot-password", controller.ForgotPasswordHandler)
		// r.Get("/reset-password-email", controller.ResetPasswordEmail)
		// r.Post("/reset-password", controller.ResetPasswordHandler)
	})



	// User
	r.Route("/user", func(r chi.Router) {
		r.Get("/", controller.GetUser)

		
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", controller.GetUserById)
			r.Patch("/", controller.UpdateUser)
			r.Delete("/", controller.DeleteUser)
		})
			
	})

	// Election
	r.Route("/election", func(r chi.Router) {
		r.Post("/createElection", controller.CreateElection)
		r.Get("/", controller.GetElection)

		r.Route("/{electionId}", func(r chi.Router) {
			r.Get("/", controller.GetElectionById)
			r.Delete("/", controller.DeleteElection)
		})
	})



	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	http.ListenAndServe(port, r)
}