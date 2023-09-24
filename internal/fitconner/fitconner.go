package fitconner

type FitConner struct {
	Name  string
	Goal1 Goals
	Goal2 Goals
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
