package models

import (
    "time"
)

type Team struct {
    ID          int               `json:"id"`
    Name        string            `json:"name"`
    ShortName   string            `json:"shortName"`
    Tla         string            `json:"tla"`
    Crest       string            `json:"crest"`
    LeagueRank  *int              `json:"leagueRank"`
    Formation   string            `json:"formation"`
    Statistics  map[string]int    `json:"statistics,omitempty"`
}

type FullTime struct {
    Home int `json:"home"`
    Away int `json:"away"`
}

type Score struct {
    Winner   string   `json:"winner"`
    Duration string   `json:"duration"`
    FullTime FullTime `json:"fullTime"`
    HalfTime FullTime `json:"halfTime"`
}

type Match struct {
    ID         int       `json:"id"`
    UTCDate    time.Time `json:"utcDate"`
    Status     string    `json:"status"`
    Minute     int       `json:"minute"`
    InjuryTime int       `json:"injuryTime"`
    Attendance int       `json:"attendance"`
    Venue      string    `json:"venue"`
    Matchday   int       `json:"matchday"`
    Stage      string    `json:"stage"`
    Group      *string   `json:"group"`
    LastUpdated time.Time `json:"lastUpdated"`
    HomeTeam   Team      `json:"homeTeam"`
    AwayTeam   Team      `json:"awayTeam"`
    Score      Score     `json:"score"`
}
