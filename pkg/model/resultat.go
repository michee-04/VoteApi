package model


type Resultat struct {
	CandidatId   string `json:"candidatId"`
	CandidatName string `json:"candidatName"`
	VoteCount    int    `json:"voteCount"`
}


func GetResultats() []Resultat {

	var resultats []Resultat
	DB.Table("votes").
        Select("candidatId, count(candidatId) as voteCount").
        Group("candidatId").
        Scan(&resultats)

	for i, r := range resultats {
		var candidat Candidat
		DB.Where("candidatId = ?", r.CandidatId).First(&candidat)
		resultats[i].CandidatName = candidat.Name
	}
	
	return resultats
}
