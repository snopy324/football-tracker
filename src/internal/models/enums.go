package models

type CompetitionID int

const (
	UEFAChampionsLeague CompetitionID = 2001
	PremierLeague       CompetitionID = 2021
	Ligue1              CompetitionID = 2015
	Bundesliga          CompetitionID = 2002
	SerieA              CompetitionID = 2019
	LaLiga              CompetitionID = 2014
)

var AllCompetitions = []CompetitionID{
	UEFAChampionsLeague,
	PremierLeague,
	Ligue1,
	Bundesliga,
	SerieA,
	LaLiga,
}
