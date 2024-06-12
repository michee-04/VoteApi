package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
)


func CreateElection(w http.ResponseWriter, r *http.Request){
	e := &model.Election{}
	utils.ParseBody(r, e)
	election := e.CreateElection()
	res, _ := json.Marshal(election)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetElection(w http.ResponseWriter, r *http.Request){
	election:= model.GetAllElection()
	res, _ := json.Marshal(election)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetElectionById(w http.ResponseWriter, r *http.Request){
	electionId := chi.URLParam(r, "electionId")
	election, _ := model.GetElectionById(electionId)

	if election == nil {
		http.Error(w, "Election not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(election)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteElection(w http.ResponseWriter, r *http.Request){
	electionId := chi.URLParam(r, "electionId")
	election := model.DeleteElection(electionId)
	res, _ := json.Marshal(election)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}