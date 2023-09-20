package main

type FitConner struct {
	Name  string
	Meta1 Metas
	Meta2 Metas
}

type Metas struct {
	FatPercentage string
	LeanMass      string
	VisceralFat   string
}

func (m *Metas) New(fatPercentage *string, leanMass *string, visceralFat *string) Metas {
	var fp, lm, vf string

	fp = *fatPercentage
	lm = *leanMass
	vf = *visceralFat

	if fatPercentage == nil {
		fp = "*"
	}
	if leanMass == nil {
		lm = "*"
	}
	if visceralFat == nil {
		vf = "*"
	}

	return Metas{
		FatPercentage: fp,
		LeanMass:      lm,
		VisceralFat:   vf,
	}
}
