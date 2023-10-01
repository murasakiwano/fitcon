package main

type FitConner struct {
	Goal1      Goals
	Goal2      Goals
	TeamName   string
	Name       string
	TeamNumber int
	Register   string
}

type Goals struct {
	FatPercentage string
	LeanMass      string
	VisceralFat   string
}

func BuildGoal1(fatPercentage string, leanMass string) Goals {
	return BuildGoal2(fatPercentage, leanMass, "")
}

func BuildGoal2(fatPercentage string, leanMass string, visceralFat string) Goals {
	var fp, lm, vf string

	fp = fatPercentage
	lm = leanMass
	vf = visceralFat

	if fatPercentage == "" {
		fp = "-"
	}
	if leanMass == "" {
		lm = "-"
	}
	if visceralFat == "" {
		vf = "-"
	}

	return Goals{
		FatPercentage: fp,
		LeanMass:      lm,
		VisceralFat:   vf,
	}
}
