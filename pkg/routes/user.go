package routes

import (
	"github.com/go-chi/chi"
	"github.com/michee/micgram/pkg/controller"
)


func RegisterUser(route *chi.Mux){

	route.HandleFunc("/auth/register", controller.CreateUser)

}