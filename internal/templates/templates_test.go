package templates

import (
	"bytes"
	"testing"

	"github.com/approvals/go-approval-tests"
	"github.com/murasakiwano/fitcon_templates/internal/fitconner"
)

func TestReplacesAFitConner(t *testing.T) {
	goal1 := fitconner.BuildGoal1("Diminuir 8%", "Aumentar 1 kg")
	goal2 := fitconner.BuildGoal2("Diminuir 4%", "Aumentar 5 kg", "")
	fitConner := fitconner.FitConner{
		Name:  "Monkey D. Luffy",
		Goal1: goal1,
		Goal2: goal2,
	}

	t.Run("it parses a FitConner into an HTML table", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := ParseFitConnerTable(&buf, fitConner); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
