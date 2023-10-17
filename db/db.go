package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
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

// creates database and returns a new DB
func New(logger *zap.SugaredLogger) (*DB, error) {
	db, err := sqlx.Connect("sqlite3", os.Getenv("DATABASE_FILE"))
	if err != nil {
		logger.Errorw("Error connecting to the database", zap.String("dbFile", os.Getenv("DATABASE_URL")), zap.Error(err))
	}

	database := DB{db: db, logger: logger}
	database.Create()
	logger.Debugw("Database created", zap.String("name", DbName))

	return &database, nil
}

func (db *DB) Create() {
	// exec the schema or fail;
	db.db.MustExec(fitconnerSchema)
}

func (db *DB) Drop() {
	db.db.MustExec(fitconnerDrop)
}
