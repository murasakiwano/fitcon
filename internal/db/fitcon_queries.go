package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murasakiwano/fitcon/internal/fitconner"
	"go.uber.org/zap"
)

// get a fitconner by id (matricula)
func (db *DB) GetFitConner(id string) (*fitconner.FitConner, error) {
	var fitConner fitconner.FitConner

	db.logger.Debugw("Id is", zap.String("matricula", id))
	if err := db.ValidateId(id); err != nil {
		db.logger.Errorw("Error while validating id", zap.Error(err))
		return nil, err
	}

	err := db.db.Get(&fitConner, getQuery, strings.ToUpper(id))
	if err != nil {
		db.logger.Errorw("Error while getting fitconner", zap.String("fitConnerId", id), zap.Error(err))
		return nil, err
	}

	db.logger.Debugw("FitConner found", zap.Object("fitconner", fitConner))
	return &fitConner, nil
}

// insert a fitconner into the database
func (db *DB) CreateFitConner(fc fitconner.FitConner) error {
	db.logger.Debugw("Trying to create fitconner", zap.Object("fitconner", fc))
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

func (db *DB) GetAllFitConners() {
	fcs := make([]fitconner.FitConner, 0)
	db.db.Select(&fcs, "SELECT * FROM fitconners;")

	db.logger.Debugw("FitConners found", zap.Objects[fitconner.FitConner]("fitconners", fcs))
}

func (db *DB) DeleteFitConner(id string) error {
	fc, err := db.GetFitConner(id)
	if err != nil {
		db.logger.Error("Error retrieving fitconner", zap.String("id", id), zap.Error(err))
		return err
	}

	db.db.MustExec("DELETE FROM fitconners WHERE id=?", fc.ID)

	return nil
}

// WARN: CURRENTLY THIS IS EXTREMELY INEFFICIENT
func (db *DB) UpdateFitConner(id string, fields map[string]string) (fitconner.FitConner, error) {
	db.logger.Debugw("Updating fitconner", zap.String("id", id), zap.Any("fields", fields))
	tx := db.db.MustBegin()

	for k, v := range fields {
		db.logger.Debugw("Updating field", zap.String("field", k), zap.String("value", v))
		tx.MustExec(
			fmt.Sprintf(updateFitConnerQuery, FitConnersTable, k),
			v,
			id,
		)
	}

	tx.Commit()

	fc, err := db.GetFitConner(id)
	if err != nil {
		return fitconner.FitConner{}, err
	}

	return *fc, nil
}
