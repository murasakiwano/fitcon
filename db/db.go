package db

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/murasakiwano/fitcon/fitconner"
	"go.uber.org/zap"
)

var (
	fitconnerSchema = fmt.Sprintf(`
create table if not exists %s (
	id varchar(7) primary key,
	name text not null,
	team_name text not null,
	team_number integer not null,
	goal1_fat_percentage text default '-',
	goal1_lean_mass text default '-',
	goal2_fat_percentage text default '-',
	goal2_lean_mass text default '-',
	goal2_visceral_fat text default '-'
);`, FitConnersTable)
	fitconnerDrop        = fmt.Sprintf("drop table %s;", FitConnersTable)
	getQuery             = fmt.Sprintf("SELECT * FROM %s WHERE id=$1", FitConnersTable)
	insertFitConnerQuery = fmt.Sprintf(`
INSERT INTO %s (
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
);`, FitConnersTable)
)

type DB struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

type ValidationError struct {
	Message string
	Code    int
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s", ve.Message)
}

var (
	INVALID_LENGTH          = &ValidationError{Message: "id must have 7 characters", Code: 0}
	INVALID_FIRST_CHARACTER = &ValidationError{Message: "id must start with 'C' or 'c'", Code: 1}
	PARSE_ERROR             = &ValidationError{Message: "could not convert %v into a number", Code: 2}
)

// creates database and returns a new DB
func New(logger *zap.SugaredLogger, dbName string) (*DB, error) {
	db, err := sqlx.Connect("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Errorw("Error connecting to the database", zap.Error(err))
	}

	database := DB{db: db, logger: logger}
	database.Create()
	logger.Debugw("Database created", zap.String("name", dbName))

	return &database, nil
}

func (db *DB) Create() {
	// exec the schema or fail;
	db.db.MustExec(fitconnerSchema)
}

func (db *DB) Drop() {
	db.db.MustExec(fitconnerDrop)
}

// get a fitconner by id (matricula)
func (db *DB) GetFitConner(id string) (*fitconner.FitConner, error) {
	var fitconner fitconner.FitConner

	db.logger.Debugw("Id is", zap.String("matricula", id))
	if err := db.ValidateId(id); err != nil {
		db.logger.Errorw("Error while validating id", zap.Error(err))
		return nil, err
	}

	err := db.db.Get(&fitconner, getQuery, strings.ToUpper(id))
	if err != nil {
		db.logger.Errorw("Error while getting fitconner", zap.String("fitConnerId", id), zap.Error(err))
		return nil, err
	}

	db.logger.Debugw("FitConner found", zap.String("id", fitconner.ID))
	return &fitconner, nil
}

// insert a fitconner into the database
func (db *DB) CreateFitConner(fc fitconner.FitConner) error {
	if err := db.ValidateId(fc.ID); err != nil {
		db.logger.Errorw("Error while validating fitconner", zap.Error(err))
		return err
	}
	_, err := db.db.NamedExec(insertFitConnerQuery, fc)
	if err != nil {
		db.logger.Errorw("Error while creating fitconner", zap.Error(err))
		return err
	}

	db.logger.Infow("FitConner created", zap.String("id", fc.ID))
	return nil
}

// batch insert fitconners into the database
func (db *DB) BatchInsert(fcs []fitconner.FitConner) error {
	for _, fc := range fcs {
		if err := db.ValidateId(fc.ID); err != nil {
			db.logger.Errorw("Error while validating fitconner", zap.Error(err))
			return err
		}
		db.logger.Debugw("insert query", zap.String("query", insertFitConnerQuery))
		db.logger.Debugw("FitConner validated", zap.Object("fitconner", fc))
		_, err := db.db.NamedExec(insertFitConnerQuery, fc)
		if err != nil {
			db.logger.Errorw("Error while batch inserting fitconners", zap.Error(err))
			return err
		}
	}

	db.logger.Info("Successfully batch-inserted fitconners")

	return nil
}

func (db *DB) ValidateId(id string) error {
	if len(id) != 7 {
		db.logger.Errorw("Invalid length", zap.String("id", id), zap.Int("id len", len(id)))
		return INVALID_LENGTH
	}

	if id[0] != 'C' && id[0] != 'c' {
		db.logger.Errorw("Invalid first character", zap.String("id", id))
		return INVALID_FIRST_CHARACTER
	}

	if _, err := strconv.Atoi(id[1:]); err != nil {
		db.logger.Errorw("Error while parsing id", zap.Error(err))
		return PARSE_ERROR
	}

	return nil
}

func (db *DB) CloseDB() {
	db.db.Close()
}

func (db *DB) GetAllFitConners(fcs *[]fitconner.FitConner) {
	db.db.Select(&fcs, "SELECT * FROM fitconners;")

	db.logger.Debugw("FitConners found", zap.Objects[fitconner.FitConner]("fitconners", *fcs))
}
