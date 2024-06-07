package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michee/micgram/pkg/routes"
)

const port = ":3000"

func main (){
	r := chi.NewRouter()

	routes.RegisterUser(r)

	http.Handle("/", r)

	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	http.ListenAndServe(port, r)
}