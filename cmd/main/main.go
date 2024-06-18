package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/michee/micgram/pkg/access"
	"github.com/michee/micgram/pkg/controller"
)

const port = ":3000"

var tokenAuth *jwtauth.JWTAuth

func main() {
	tokenAuth = jwtauth.New("HS256", []byte("ksQD5adHXZ-5SSJCupcHwBzDi6q5kfr5hdU7Eq5tMmo"), nil)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
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

	r.Group(func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		// Election
		r.Route("/election", func(r chi.Router) {
			r.With(access.AdminOnly).Post("/create", controller.CreateElection)
			r.Get("/", controller.GetElection)
			r.Route("/{electionId}", func(r chi.Router) {
				r.Get("/", controller.GetElectionById)
				r.With(access.AdminOnly).Patch("/", controller.UpdateElection)
				r.With(access.AdminOnly).Delete("/", controller.DeleteElection)
			})
		})

		// Candidat
		r.Route("/candidat", func(r chi.Router) {
			r.With(access.AdminOnly).Post("/create", controller.CreateCaddidat)
			r.Get("/", controller.GetCandidat)

			r.Route("/{candidatId}", func(r chi.Router) {
				r.Get("/", controller.GetCandidatById)
				r.With(access.AdminOnly).Patch("/", controller.UpdateCandidat)
				r.With(access.AdminOnly).Delete("/", controller.DeleteCandidat)
			})
		})

	})

	// Vote
	r.Route("/vote", func(r chi.Router) {
		r.Post("/create/{userId}/{electionId}/{candidatId}", controller.CreateVote)
		r.Get("/", controller.GetVote)
		r.Route("/{voteId}", func(r chi.Router) {
			r.Get("/", controller.GetVoteByid)
			r.With(access.AdminOnly).Delete("/", controller.DeleteVote)
		})
	})

	r.Get("/resultats", controller.GetResultats)

	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	http.ListenAndServe(port, r)
}
