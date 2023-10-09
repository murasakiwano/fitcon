package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/murasakiwano/fitcon/fitconner"
)

func ParseConsolidations(filename string) ([]byte, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	r := csv.NewReader(fileReader)
	records, err := r.ReadAll()
	if err != nil {
		return []byte{}, err
	}

	categories := records[0]
	participants := records[1:]
	fmt.Println(categories)
	fmt.Println(participants)
	players := []fitconner.Fitconner{}

	for _, p := range participants {
		teamNumber, err := strconv.Atoi(p[0])
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
		teamName := p[1]
		register := p[2]
		name := p[3]
		goal1FatPercentage := p[4]
		goal1LeanMass := p[5]
		goal2VisceralFat := p[6]
		goal2FatPercentage := p[7]
		goal2LeanMass := p[8]
		players = append(players, fitconner.Fitconner{
			TeamNumber: teamNumber,
			TeamName:   teamName,
			Register:   register,
			Goal1:      fitconner.BuildGoal1(goal1FatPercentage, goal1LeanMass),
			Goal2:      fitconner.BuildGoal2(goal2FatPercentage, goal2LeanMass, goal2VisceralFat),
			Name:       name,
		})
	}

	for _, p := range players {
		fmt.Println(p.Name, p.Goal1, p.Goal2, p.TeamName, p.Register)
	}

	return nil, nil
}
