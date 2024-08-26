package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"football_tracker/src/internal/api"
	"football_tracker/src/internal/config"
	"football_tracker/src/internal/models"
)

type FootballService struct {
	footballClient *api.FootBallClient
	lineClient     *api.LineMessageClient
}

func NewFootballService() *FootballService {
	return &FootballService{
		footballClient: api.NewClient(),
		lineClient:     api.NewLineClient(),
	}
}

func (s *FootballService) fetchCompetitionMatches(leagueID models.CompetitionID) (models.CompetitionWithMatch, error) {
	return s.footballClient.FetchCompetitionMatches(leagueID)
}

func findTeamByID(teams []models.Team, id int) models.Team {
	for _, team := range teams {
		if team.ID == id {
			return team
		}
	}
	return models.Team{}
}

func (s *FootballService) BrocastRecentCompletedMatches() {

	// ========================================
	// for each models.CompetitionID, fetch the recent matches
	for _, competitionID := range models.AllCompetitions {

		competition, err := s.footballClient.FetchCompetitionMatches(competitionID)

		if err != nil {
			log.Printf("Error fetching competition matches: %v", err)
			return
		}

		var recentMatches []models.Match
		now := time.Now()
		twentyFourHoursAgo := now.Add(-24 * time.Hour)

		for _, match := range competition.Matches {
			if match.Status == "FINISHED" && match.UTCDate.After(twentyFourHoursAgo) && match.UTCDate.Before(now) {
				recentMatches = append(recentMatches, match)
			}
		}

		// ========================================
		// Prepare message for Line
		var message models.LinePayload = models.LinePayload{
			To: config.LineToID,
			Messages: []models.Message{
				{
					Type:    "flex",
					AltText: "Recent matches",
					Contents: models.Content{
						Type: "bubble",
						Size: "giga",
						Body: &models.Body{
							Type:       "box",
							Layout:     "vertical",
							Contents:   make([]models.Content, 0),
							Spacing:    "md",
							PaddingAll: "12px",
						},
					},
				},
			},
		}

		var body = &message.Messages[0].Contents.Body

		matchBox := models.Content{
			Type:       "box",
			Layout:     "horizontal",
			PaddingAll: "8px",
			Contents: []models.Content{
				{
					Type:   "text",
					Text:   competition.Competition.Name,
					Size:   "lg",
					Align:  "center",
					Wrap:   true,
					Flex:   12,
					Weight: "bold",
				},
			},
		}

		(*body).Contents = append((*body).Contents, matchBox)

		for _, match := range recentMatches {

			matchBox := models.Content{
				Type:       "box",
				Layout:     "horizontal",
				PaddingAll: "8px",
				Contents: []models.Content{
					{
						Type:   "box",
						Layout: "horizontal",
						Contents: []models.Content{
							{
								Type:       "image",
								URL:        fmt.Sprintf("https://crests.football-data.org/%d.svg", match.HomeTeam.ID),
								Size:       "xxs",
								AspectMode: "fit",
								Flex:       1,
							},
							{
								Type:    "text",
								Text:    match.HomeTeam.Name,
								Size:    "xs",
								Align:   "start",
								Wrap:    true,
								Flex:    5,
								Weight:  "bold",
								Gravity: "center",
								Margin:  "md",
							},
						},
						Flex: 4,
					},
					{
						Type:   "text",
						Text:   fmt.Sprintf("%d - %d", match.Score.FullTime.Home, match.Score.FullTime.Away),
						Size:   "sm",
						Align:  "center",
						Wrap:   false,
						Flex:   2,
						Weight: "bold",
					},
					{
						Type:   "box",
						Layout: "horizontal",
						Contents: []models.Content{
							{
								Type:    "text",
								Text:    match.AwayTeam.Name,
								Size:    "xs",
								Align:   "end",
								Wrap:    true,
								Flex:    5,
								Weight:  "bold",
								Gravity: "center",
							},
							{
								Type:       "image",
								URL:        fmt.Sprintf("https://crests.football-data.org/%d.svg", match.AwayTeam.ID),
								Size:       "xxs",
								AspectMode: "fit",
								Flex:       1,
								Margin:     "md",
							},
						},
						Flex: 4,
					},
				},
			}
			(*body).Contents = append((*body).Contents, matchBox)
		}

		// ========================================
		// Send message to Line
		err = s.lineClient.PushMessage(message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}

		// ========================================
		// Print to console
		if len(recentMatches) == 0 {
			fmt.Printf("\n\nNo matches for %s the last 24 hours.\n\n", competition.Competition.Name)
		} else {
			fmt.Printf("\n\nMatches completed in the last 24 hours for %s:\n\n", competition.Competition.Name)

			// 找出最長的球隊名稱長度
			maxTeamNameLength := 0
			for _, match := range recentMatches {
				if len(match.HomeTeam.Name) > maxTeamNameLength {
					maxTeamNameLength = len(match.HomeTeam.Name)
				}
				if len(match.AwayTeam.Name) > maxTeamNameLength {
					maxTeamNameLength = len(match.AwayTeam.Name)
				}
			}

			// 打印每場比賽
			for _, match := range recentMatches {
				homeTeam := padRight(match.HomeTeam.Name, maxTeamNameLength)
				awayTeam := padRight(match.AwayTeam.Name, maxTeamNameLength)

				fmt.Printf("%s %2d - %-2d %s  (%s)\n",
					homeTeam,
					match.Score.FullTime.Home,
					match.Score.FullTime.Away,
					awayTeam,
					match.UTCDate.Format("2006-01-02 15:04 MST"))
			}
		}
	}
}

func padRight(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}
