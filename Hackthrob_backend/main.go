package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Hackathon struct {
	HackathonTitle string `json:"hackathontitle"`
	Prizes string `json:"prizes"`
	Achievements string `json:"achievements"`
	Eligibility string `json:"eligibility"`
	Requirements string `json:"requirements"`
	JudgingCriteria string `json:"judgingcriteria"`
}

// Return type for hackathons
type Hackathons []Hackathon

func allHackathons(w http.ResponseWriter, r*http.Request) {
	hackathons := Hackathons {
		Hackathon{
			HackathonTitle:"Hack 1",
			Prizes: "Prize 1",
			Achievements: "Achievement 1",
			Eligibility: "Eligibility 1",
			Requirements: "Requirement 1",
			JudgingCriteria: "Judging Criteria 1",
		},
		Hackathon{
			HackathonTitle:"Hack 2",
			Prizes: "Prize 2",
			Achievements: "Achievement 2",
			Eligibility: "Eligibility 2",
			Requirements: "Requirement 2",
			JudgingCriteria: "Judging Criteria 2",
		},
	}

	fmt.Println("All Hackathons - sourced")
	json.NewEncoder(w).Encode(hackathons)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200 OK")
}

func handleRequests() {
	http.HandleFunc("/healthcheck", healthCheck)
	http.HandleFunc("/getAllHackathons", allHackathons)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
