package entity

import (
	"strconv"
	"time"
)

type MatchResult struct {
	teamAScore int
	teamBScore int
}

func NewMatchResult(teamAScore, teamBScore int) *MatchResult {
	return &MatchResult{
		teamAScore: teamAScore,
		teamBScore: teamBScore,
	}
	
}

func (m *MatchResult) GetResult() string {
	return strconv.Itoa(m.teamAScore) + "-" + strconv.Itoa(m.teamBScore)

}

type Match struct {
	ID       string
	TeamA    *Team
	TeamB    *Team
	TeamA_ID string
	TeamB_ID string
	Date     time.Time
	Status   string
	Actions []GameAction
	Result   MatchResult
}

func NewMatch(id string, teamA *Team, teamB *Team, date time.Time) *Match{
	return &Match{
		ID: id,
		TeamA: teamA,
		TeamB: teamB,
		TeamA_ID: teamA.ID,
		TeamB_ID: teamB.ID,
		Date: date,
	}
}
