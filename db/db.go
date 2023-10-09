package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/murasakiwano/fitcon/fitconner"
	"go.uber.org/zap"
)

var fitconnerSchema = `
create table if not exists fitcon_metas (
	id varchar(7) primary key,
	name text not null,
	team_name text not null,
	team_number integer not null,
	goal1_fat_percentage text default '-',
	goal1_lean_mass text default '-',
	goal2_fat_percentage text default '-',
	goal2_lean_mass text default '-',
	goal2_visceral_fat text default '-'
);`

var fitconnerDrop = "drop table fitcon_metas;"

var insertFitconnerQuery = `
INSERT INTO fitcon_metas (
	id,
	name,
	team_name,
	team_number,
	goal1_fat_percentage,
	goal1_lean_mass,
	goal2_fat_percentage,
	goal2_lean_mass,
	goal2_visceral_fat
)
VALUES (
	:id,
	:name,
	:team_name,
	:team_number,
	:goal1_fat_percentage,
	:goal1_lean_mass,
	:goal2_fat_percentage,
	:goal2_lean_mass,
	:goal2_visceral_fat
)`

type DB struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

// creates database and returns a new DB
func New(logger zap.SugaredLogger) (*DB, error) {
	// this Pings the database trying to connect
	db, err := sqlx.Connect("sqlite3", DbName)
	if err != nil {
		logger.Error("Error connecting to the database", zap.Error(err))
	}

	// exec the schema or fail;
	db.MustExec(fitconnerSchema)
	if err != nil {
		logger.Error("Error while creating table", zap.Error(err))
		return nil, err
	}

	return &DB{db: db, logger: &logger}, nil
}

// get a fitconner by id (matricula)
func (db *DB) GetFitConner(id string) (*fitconner.Fitconner, error) {
	var fitconner fitconner.Fitconner
	err := db.db.Get(&fitconner, "SELECT * FROM fitcon_metas WHERE id=$1", id)
	if err != nil {
		db.logger.Error("Error while getting fitconner", zap.Error(err))
		return nil, err
	}

	db.logger.Debugw("Fitconner found", zap.String("id", fitconner.ID))
	return &fitconner, nil
}

// insert a fitconner into the database
func (db *DB) CreateFitConner(fc fitconner.Fitconner) error {
	_, err := db.db.NamedExec(insertFitconnerQuery, fc)
	if err != nil {
		db.logger.Error("Error while creating fitconner", zap.Error(err))
		return err
	}

	db.logger.Infow("Fitconner created", zap.String("id", fc.ID))
	return nil
}

// batch insert fitconners into the database
func (db *DB) BatchInsert(fcs []fitconner.Fitconner) error {
	_, err := db.db.NamedExec(insertFitconnerQuery, fcs)
	if err != nil {
		db.logger.Error("Error while batch inserting fitconners", zap.Error(err))
		return err
	}

	db.logger.Info("Successfully batch-inserted fitconners")

	return nil
}
