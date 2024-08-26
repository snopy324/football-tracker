package models

import (
	"time"
)

type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CompetitionInfo struct {
	ID          int       `json:"id"`
	Area        Area      `json:"area"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Plan        string    `json:"plan"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type CompetitionWithMatch struct {
	Matches     []Match         `json:"matches"`
	Competition CompetitionInfo `json:"competition"`
}

type CompetitionWithTeam struct {
	Teams       []Team          `json:"teams"`
	Competition CompetitionInfo `json:"competition"`
}
