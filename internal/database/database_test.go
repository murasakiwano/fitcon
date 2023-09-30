package database_test

import (
	"testing"

	"github.com/murasakiwano/fitcon_templates/internal/database"
	"github.com/murasakiwano/fitcon_templates/internal/fitconner"
)

func TestDatabase(t *testing.T) {
	db, err := database.InitDb("./fitcon_test.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	participant := fitconner.FitConner{
		Name: "test",
		Goal1: fitconner.BuildGoal1(
			"10",
			"20",
		),
		Goal2: fitconner.BuildGoal2(
			"10",
			"20",
			"30",
		),
	}

	t.Run("insert users into table", func(t *testing.T) {
		err = database.InsertFitconnerIntoTable(db, participant)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("retrieve users from table", func(t *testing.T) {
		err = database.InsertFitconnerIntoTable(db, participant)
		if err != nil {
			t.Error(err)
		}

		expectedName := participant.Name
		got, err := database.RetrieveFitconnerFromTable(db, expectedName)
		if err != nil {
			t.Error(err)
		}
		if got.Name != expectedName {
			t.Errorf("Expected %s, got %s", expectedName, got.Name)
		}
	})

	_, err = db.Exec("delete from fitcon_metas;")
	if err != nil {
		t.Error(err)
	}
}
