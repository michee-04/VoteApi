package model

import "log"

type Resultat struct {
	CandidatId   string `json:"candidatId"`
	CandidatName string `json:"candidatName"`
	VoteCount    int    `json:"voteCount"`
}


func GetResultats() []Resultat {
	var resultats []Resultat

	rows, err := DB.Table("votes").
			Select("candidat_id, COUNT(*) as voteCount").
			Where("status = ?", true).
			Group("candidat_id").
			Order("voteCount DESC").
			Rows()
	if err != nil {
			log.Fatalf("Error querying votes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
			var resultat Resultat
			if err := rows.Scan(&resultat.CandidatId, &resultat.VoteCount); err != nil {
					log.Fatalf("Error scanning row: %v", err)
			}
			var candidat Candidat
			if err := DB.Where("CandidatId = ?", resultat.CandidatId).First(&candidat).Error; err != nil {
					log.Printf("Error fetching candidat: %v", err)
			} else {
					resultat.CandidatName = candidat.Name
			}
			resultats = append(resultats, resultat)
	}

	return resultats
}
