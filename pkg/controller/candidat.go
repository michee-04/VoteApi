package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
)


func CreateCaddidat(w http.ResponseWriter, r *http.Request){
	candidat := &model.Candidat{}
	utils.ParseBody(r, candidat)
	c := candidat.CreateCandidat()
	res, _ := json.Marshal(c)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCandidat(w http.ResponseWriter, r *http.Request) {
	c := model.GetCandidat()
	res, _ := json.Marshal(c)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetCandidatById(w http.ResponseWriter, r *http.Request){
	candidatId := chi.URLParam(r, "candidatId")
	c, _ := model.GetCandidatById(candidatId)

	if c == nil {
		http.Error(w, "Candidat not found", http.StatusNotFound)
	}
	res, _ := json.Marshal(c)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateCandidat(w http.ResponseWriter, r *http.Request) {
	candidatUpdate := &model.Candidat{}
	utils.ParseBody(r, &candidatUpdate)
	candidatId := chi.URLParam(r, "candidatId")
	candidatDetail, db := model.GetCandidatById(candidatId)

	if candidatUpdate.Name != ""{
		candidatDetail.Name = candidatUpdate.Name
	}
	if candidatUpdate.Email != ""{
		candidatDetail.Email = candidatUpdate.Email
	}
	if candidatUpdate.Discour != ""{
		candidatDetail.Discour = candidatUpdate.Discour
	}

	db.Save(&candidatDetail)

	res, _ := json.Marshal(candidatDetail)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteCandidat(w http.ResponseWriter, r *http.Request) {
	candidatId := chi.URLParam(r, "candidatId")
	c := model.DeleteCandidat(candidatId)
	res, _ := json.Marshal(c)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}