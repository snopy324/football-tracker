package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"football_tracker/src/internal/config"
	"football_tracker/src/internal/models"
)

type FootBallClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient() *FootBallClient {
	return &FootBallClient{
		baseURL:    config.FootballDataBaseURL,
		apiKey:     config.FootballDataAPIKey,
		httpClient: &http.Client{},
	}
}

func (c *FootBallClient) FetchTeams(competitionID models.CompetitionID) (models.CompetitionWithTeam, error) {
	url := fmt.Sprintf("%s/competitions/%d/teams", c.baseURL, competitionID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.CompetitionWithTeam{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-Auth-Token", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.CompetitionWithTeam{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CompetitionWithTeam{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var competition models.CompetitionWithTeam
	err = json.NewDecoder(resp.Body).Decode(&competition)
	if err != nil {
		return models.CompetitionWithTeam{}, fmt.Errorf("error decoding response: %w", err)
	}

	return competition, nil
}

func (c *FootBallClient) FetchCompetitionMatches(competitionID models.CompetitionID) (models.CompetitionWithMatch, error) {
	url := fmt.Sprintf("%s/competitions/%d/matches", c.baseURL, competitionID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.CompetitionWithMatch{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-Auth-Token", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.CompetitionWithMatch{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CompetitionWithMatch{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var competition models.CompetitionWithMatch
	err = json.NewDecoder(resp.Body).Decode(&competition)
	if err != nil {
		return models.CompetitionWithMatch{}, fmt.Errorf("error decoding response: %w", err)
	}

	return competition, nil
}
