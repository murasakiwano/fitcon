package fitconner

type Fitconner struct {
	TeamName           string `json:"teamName"  form:"teamName" db:"team_name"`
	Name               string `json:"name"  form:"name" db:"name"`
	ID                 string `json:"matricula"  form:"matricula" db:"id"`
	Goal1FatPercentage string `json:"goal1FatPercentage"  form:"goal1FatPercentage" db:"goal1_fat_percentage"`
	Goal1LeanMass      string `json:"goal1LeanMass"  form:"goal1LeanMass" db:"goal1_lean_mass"`
	Goal2FatPercentage string `json:"goal2FatPercentage"  form:"goal2FatPercentage" db:"goal2_fat_percentage"`
	Goal2LeanMass      string `json:"goal2LeanMass"  form:"goal2LeanMass" db:"goal2_lean_mass"`
	Goal2VisceralFat   string `json:"goal2VisceralFat"  form:"goal2VisceralFat" db:"goal2_visceral_fat"`
	TeamNumber         int    `json:"teamNumber"  form:"teamNumber" db:"team_number"`
}

func New(
	id,
	name,
	teamName,
	g1FatPercentage,
	g1LeanMass,
	g2FatPercentage,
	g2LeanMass,
	g2VisceralFat string,
	teamNumber int,
) *Fitconner {
	goals1 := buildGoals(g1FatPercentage, g1LeanMass, "")
	goals2 := buildGoals(g2FatPercentage, g2LeanMass, g2VisceralFat)
	return &Fitconner{
		ID:                 id,
		Name:               name,
		TeamName:           teamName,
		TeamNumber:         teamNumber,
		Goal1FatPercentage: goals1["fat_percentage"],
		Goal1LeanMass:      goals1["lean_mass"],
		Goal2FatPercentage: goals2["fat_percentage"],
		Goal2LeanMass:      goals2["lean_mass"],
		Goal2VisceralFat:   goals2["visceral_fat"],
	}
}

func buildGoals(fatPercentage string, leanMass string, visceralFat string) map[string]string {
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

	return map[string]string{
		"fatPercentage": fp,
		"leanMass":      lm,
		"visceralFat":   vf,
	}
}
