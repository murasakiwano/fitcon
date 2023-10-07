package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/murasakiwano/fitcon/fitconner"
)

type FitConnerStore struct {
	db        *sql.DB
	tableName string
}

// creates database and returns a new FitConnerStore
func NewFitConnerStore(dbName string) (*FitConnerStore, error) {
	createDbQuery := `
create table if not exists fitcon_metas (
	id varchar(7) primary key,
	name text not null,
	goal1_fat_percentage text default '-',
	goal1_lean_mass text default '-',
	goal2_fat_percentage text default '-',
	goal2_lean_mass text default '-',
	goal2_visceral_fat text default '-'
);
delete from fitcon_metas;
`
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		slog.Error("%v", err)
		return nil, err
	}
	_, err = db.Exec(createDbQuery)
	if err != nil {
		slog.Error("%v", err)
		return nil, err
	}

	return &FitConnerStore{db: db, tableName: "fitcon_metas"}, nil
}

func (fcs *FitConnerStore) InsertFitconner(participant fitconner.FitConner) error {
	query := `
	insert into fitcon_metas (id, name, goal1_fat_percentage, goal1_lean_mass, goal2_fat_percentage, goal2_lean_mass, goal2_visceral_fat) values (?, ?, ?, ?, ?, ?);
	`
	_, err := fcs.db.Exec(
		query,
		participant.Register,
		participant.Name,
		participant.Goal1.FatPercentage,
		participant.Goal1.LeanMass,
		participant.Goal2.FatPercentage,
		participant.Goal2.LeanMass,
		participant.Goal2.VisceralFat,
	)
	if err != nil {
		slog.Error("%v", err)
		return err
	}

	return nil
}

// Deletes by id
func (fcs *FitConnerStore) DeleteFitconner(participant fitconner.FitConner) error {
	query := "delete from %s where id = ?;"
	query = fmt.Sprintf(query, fcs.tableName)
	result, err := fcs.db.Exec(query, participant.Name)
	if err != nil {
		slog.Error("%v", err)
		return err
	}

	slog.Info("Result: %v", result)

	return nil
}

func (fcs *FitConnerStore) GetFitconner(playerId string) (*fitconner.FitConner, error) {
	query := "select * from %s where id = $1"
	query = fmt.Sprintf(query, fcs.tableName)
	result := fcs.db.QueryRow(query, playerId)

	var id,
		name,
		goal1_fat_percentage,
		goal1_lean_mass,
		goal2_fat_percentage,
		goal2_lean_mass,
		goal2_visceral_fat string

	if err := result.Scan(
		&id,
		&name,
		&goal1_fat_percentage,
		&goal1_lean_mass,
		&goal2_fat_percentage,
		&goal2_lean_mass,
		&goal2_visceral_fat); err != nil {
		slog.Error("Error encountered while querying table - %v", err)
		return nil, err
	}

	participant := fitconner.FitConner{
		Name:  name,
		Goal1: fitconner.BuildGoal1(goal1_fat_percentage, goal1_lean_mass),
		Goal2: fitconner.BuildGoal2(goal2_fat_percentage, goal2_lean_mass, goal2_visceral_fat),
	}

	slog.Info("Got participant: %+v", participant)

	return &participant, nil
}
