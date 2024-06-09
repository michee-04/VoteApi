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
	
	// http.Handle("/", r)
	r.Use(middleware.Logger)

	r.Post("/auth/register", controller.CreateUser)

	r.Get("/auth/verify", controller.VerifyHandler)
	r.Post("/auth/login", controller.LoginHandler)


	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	http.ListenAndServe(port, r)
}