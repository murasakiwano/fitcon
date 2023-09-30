package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/murasakiwano/fitcon_templates/internal/fitconner"
)

func InitDb(dbName string) (*sql.DB, error) {
	createDbQuery := `
create table if not exists fitcon_metas (
	id integer primary key autoincrement,
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
		log.Fatal(err)
		return nil, err
	}
	_, err = db.Exec(createDbQuery)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func InsertFitconnerIntoTable(db *sql.DB, participant fitconner.FitConner) error {
	query := `
	insert into fitcon_metas (name, goal1_fat_percentage, goal1_lean_mass, goal2_fat_percentage, goal2_lean_mass, goal2_visceral_fat) values (?, ?, ?, ?, ?, ?);
	`
	_, err := db.Exec(
		query,
		participant.Name,
		participant.Goal1.FatPercentage,
		participant.Goal1.LeanMass,
		participant.Goal2.FatPercentage,
		participant.Goal2.LeanMass,
		participant.Goal2.VisceralFat,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func DeleteFitconnerFromTable(db *sql.DB, participant fitconner.FitConner) error {
	query := "delete from fitcon_metas where name = ?;"
	result, err := db.Exec(query, participant.Name)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Print(result)

	return nil
}

func RetrieveFitconnerFromTable(db *sql.DB, playerName string) (*fitconner.FitConner, error) {
	query := "select * from fitcon_metas where name = $1"
	result := db.QueryRow(query, playerName)
	var id, name, goal1_fat_percentage, goal1_lean_mass, goal2_fat_percentage, goal2_lean_mass, goal2_visceral_fat string
	if err := result.Scan(&id, &name, &goal1_fat_percentage, &goal1_lean_mass, &goal2_fat_percentage, &goal2_lean_mass, &goal2_visceral_fat); err != nil {
		log.Fatalf("Error encountered while querying table - %v", err)
		return nil, err
	}

	log.Print(name)

	participant := fitconner.FitConner{
		Name:  name,
		Goal1: fitconner.BuildGoal1(goal1_fat_percentage, goal1_lean_mass),
		Goal2: fitconner.BuildGoal2(goal2_fat_percentage, goal2_lean_mass, goal2_visceral_fat),
	}

	return &participant, nil
}
