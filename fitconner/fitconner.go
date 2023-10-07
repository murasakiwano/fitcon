package fitconner

type FitConner struct {
	Goal1      Goals  `json:"goals1" xml:"goals1" form:"goals1" query:"goals1"`
	Goal2      Goals  `json:"goals2" xml:"goals2" form:"goals2" query:"goals2"`
	TeamName   string `json:"teamName" xml:"teamName" form:"teamName" query:"teamName"`
	Name       string `json:"name" xml:"name" form:"name" query:"name"`
	Register   string `json:"register" xml:"register" form:"register" query:"register"`
	TeamNumber int    `json:"teamNumber" xml:"teamNumber" form:"teamNumber" query:"teamNumber"`
}

type Goals struct {
	FatPercentage string `json:"fatPercentage" xml:"fatPercentage" form:"fatPercentage" query:"fatPercentage"`
	LeanMass      string `json:"leanMass" xml:"leanMass" form:"leanMass" query:"leanMass"`
	VisceralFat   string `json:"visceralFat" xml:"visceralFat" form:"visceralFat" query:"visceralFat"`
}

func NewFitConner(
	register,
	name,
	teamName,
	g1FatPercentage,
	g1LeanMass,
	g2FatPercentage,
	g2LeanMass,
	g2VisceralFat string,
	teamNumber int,
) *FitConner {
	g1 := BuildGoal1(g1FatPercentage, g1LeanMass)
	g2 := BuildGoal2(g2FatPercentage, g2LeanMass, g2VisceralFat)

	return &FitConner{
		Register:   register,
		Name:       name,
		TeamName:   teamName,
		TeamNumber: teamNumber,
		Goal1:      g1,
		Goal2:      g2,
	}
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
