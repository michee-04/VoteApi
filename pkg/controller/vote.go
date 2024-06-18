package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michee/micgram/pkg/model"
	"github.com/michee/micgram/pkg/utils"
)

func CreateVote(w http.ResponseWriter, r *http.Request) {
	vote := &model.Vote{}
	utils.ParseBody(r, vote)

	userId := chi.URLParam(r, "userId")
	electionId := chi.URLParam(r, "electionId")
	candidatId := chi.URLParam(r, "candidatId")

	var existingVote model.Vote
	if err := model.DB.Where("user_id = ? AND election_id = ? AND status = ?", userId, electionId, true).First(&existingVote).Error; err == nil {
		http.Error(w, "User has already voted in this election", http.StatusForbidden)
		return
	}

	v := vote.CreateVote(userId, electionId, candidatId)

	res, _ := json.Marshal(v)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetVote(w http.ResponseWriter, r *http.Request) {
	vote := model.GetVote()
	res, _ := json.Marshal(vote)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetVoteByid(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "voteId")
	v, _ := model.GetVoteById(voteId)

	res, _ := json.Marshal(v)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteVote(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "voteId")
	v := model.DeleteVote(voteId)
	res, _ := json.Marshal(v)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
