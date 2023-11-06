package fitconner

import (
	"github.com/murasakiwano/fitcon/internal/auth"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FitConner struct {
	ID                 string `json:"matricula"  form:"matricula" db:"id" query:"matricula"`
	Name               string `json:"name"  form:"name" db:"name" query:"name"`
	HashedPassword     string `json:"hashed_password"  form:"hashed_password" db:"hashed_password"`
	Goal1FatPercentage string `json:"goal1FatPercentage"  form:"goal1FatPercentage" db:"goal1_fat_percentage"`
	Goal1LeanMass      string `json:"goal1LeanMass"  form:"goal1LeanMass" db:"goal1_lean_mass"`
	Goal1Weight        string `json:"goal1Weight" form:"goal1Weight" db:"goal1_weight"`
	Goal2FatPercentage string `json:"goal2FatPercentage"  form:"goal2FatPercentage" db:"goal2_fat_percentage"`
	Goal2LeanMass      string `json:"goal2LeanMass"  form:"goal2LeanMass" db:"goal2_lean_mass"`
	Goal2VisceralFat   string `json:"goal2VisceralFat"  form:"goal2VisceralFat" db:"goal2_visceral_fat"`
	Goal2Weight        string `json:"goal2Weight" form:"goal2Weight" db:"goal2_weight"`
	TeamName           string `json:"teamName"  form:"teamName" db:"team_name"`
	TeamNumber         int    `json:"teamNumber"  form:"teamNumber" db:"team_number"`
}

var (
	logger, _ = zap.NewDevelopment()
	sugar     = logger.Sugar()
)

func New(
	id,
	name,
	password,
	teamName,
	g1FatPercentage,
	g1LeanMass,
	g1Weight,
	g2FatPercentage,
	g2LeanMass,
	g2VisceralFat,
	g2Weight string,
	teamNumber int,
) (*FitConner, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	goals1 := buildGoals(g1FatPercentage, g1LeanMass, "", g1Weight)
	goals2 := buildGoals(g2FatPercentage, g2LeanMass, g2VisceralFat, g2Weight)
	return &FitConner{
		ID:                 id,
		Name:               name,
		HashedPassword:     string(hash),
		TeamName:           teamName,
		TeamNumber:         teamNumber,
		Goal1FatPercentage: goals1["fatPercentage"],
		Goal1LeanMass:      goals1["leanMass"],
		Goal1Weight:        goals1["weight"],
		Goal2FatPercentage: goals2["fatPercentage"],
		Goal2LeanMass:      goals2["leanMass"],
		Goal2VisceralFat:   goals2["visceralFat"],
		Goal2Weight:        goals2["weight"],
	}, nil
}

func buildGoals(fatPercentage string, leanMass string, visceralFat string, weight string) map[string]string {
	var fp, lm, vf, w string

	fp = fatPercentage
	lm = leanMass
	vf = visceralFat
	w = weight

	if fatPercentage == "" {
		fp = "-"
	}
	if leanMass == "" {
		lm = "-"
	}
	if visceralFat == "" {
		vf = "-"
	}
	if w == "" {
		vf = "-"
	}

	return map[string]string{
		"fatPercentage": fp,
		"leanMass":      lm,
		"visceralFat":   vf,
		"weight":        w,
	}
}

// implement ObjectMarshaler for use with zap
func (fc FitConner) MarshalLogObject(oe zapcore.ObjectEncoder) error {
	oe.AddString("matricula", fc.ID)
	oe.AddString("teamName", fc.TeamName)
	oe.AddString("password", "<REDACTED>")
	oe.AddInt("teamNumber", fc.TeamNumber)
	oe.AddString("name", fc.Name)
	oe.AddString("goal1FatPercentage", fc.Goal1FatPercentage)
	oe.AddString("goal1LeanMass", fc.Goal1LeanMass)
	oe.AddString("goal1Weight", fc.Goal1Weight)
	oe.AddString("goal2FatPercentage", fc.Goal2FatPercentage)
	oe.AddString("goal2LeanMass", fc.Goal2LeanMass)
	oe.AddString("goal2VisceralFat", fc.Goal2VisceralFat)
	oe.AddString("goal2Weight", fc.Goal2Weight)

	return nil
}

func (fc *FitConner) ClearPassword() {
	fc.HashedPassword = ""
}

func (fc *FitConner) SetPassword(password string) error {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	fc.HashedPassword = string(hash)

	return nil
}

func (fc FitConner) PasswordEmpty() bool {
	return fc.HashedPassword == ""
}
