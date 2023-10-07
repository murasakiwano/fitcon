package db

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	fcs, err := NewFitConnerStore("./fitcon_test.db")
	if err != nil {
		t.Error(err)
	}
	defer fcs.db.Close()

	participant := FitConner{
		Name: "test",
		Goal1: BuildGoal1(
			"10",
			"20",
		),
		Goal2: BuildGoal2(
			"10",
			"20",
			"30",
		),
	}

	t.Run("insert users into table", func(t *testing.T) {
		err = fcs.InsertFitconner(participant)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("retrieve users from table", func(t *testing.T) {
		err = fcs.InsertFitconner(participant)
		if err != nil {
			t.Error(err)
		}

		expectedName := participant.Name
		got, err := fcs.GetFitconner(expectedName)
		if err != nil {
			t.Error(err)
		}
		if got.Name != expectedName {
			t.Errorf("Expected %s, got %s", expectedName, got.Name)
		}
	})

	_, err = fcs.db.Exec("delete from fitcon_metas;")
	if err != nil {
		t.Error(err)
	}
}
