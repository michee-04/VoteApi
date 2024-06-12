package model


type Resultat struct {
	CandidatId   string `json:"candidatId"`
	CandidatName string `json:"candidatName"`
	VoteCount    int    `json:"voteCount"`
}


func GetResultats() []Resultat {

	var resultats []Resultat
	DB.Table("votes").Select("candidat_id, count(candidat_id) as vote_count").Group("candidat_id").Find(&resultats)

	for i, r := range resultats {
		var candidat Candidat
		DB.Where("candidat_id = ?", r.CandidatId).First(&candidat)
		resultats[i].CandidatName = candidat.Name
	}
	
	return resultats
}
