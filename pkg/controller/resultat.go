package controller

import (
	"encoding/json"
	"net/http"

	"github.com/michee/micgram/pkg/model"
)


func GetResultats(w http.ResponseWriter, r *http.Request) {
	resultat := model.GetResultats()
	res, _ := json.Marshal(resultat)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}