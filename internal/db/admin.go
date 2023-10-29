package db

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

const (
	AdminsTable = "admins"
)

var (
	adminSchema = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	name TEXT PRIMARY KEY,
	hashed_password TEXT DEFAULT ''
);`, AdminsTable)
	dropAdmins       = fmt.Sprintf("DROP TABLE %s;", AdminsTable)
	getAdminQuery    = fmt.Sprintf("SELECT * FROM %s WHERE name=$1", AdminsTable)
	insertAdminQuery = fmt.Sprintf(`
INSERT INTO %s (
	name,
	hashed_password
)
VALUES (
	:name,
	:hashed_password
);`, AdminsTable)
	deleteAdminQuery = fmt.Sprintf("DELETE FROM %s WHERE name=?", AdminsTable)
)

type Admin struct {
	Name           string `json:"name" db:"name"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
}

func (db *DB) GetAdmin(name string) (*Admin, error) {
	var admin Admin

	db.logger.Debugw("name is", zap.String("name", name))

	err := db.db.Get(&admin, getAdminQuery, name)
	if err != nil {
		return nil, err
	}

	db.logger.Debugw("Admin found", zap.Any("admin", admin))
	return &admin, nil
}

func (db *DB) CreateAdmin(admin Admin) error {
	db.logger.Debugw("Trying to create admin", zap.Any("admin", admin))

	if admin.Name == "" {
		return errors.New("missing admin name")
	}

	if a, _ := db.GetAdmin(admin.Name); a != nil {
		return ErrAlreadyExists
	}

	_, err := db.db.NamedExec(insertAdminQuery, admin)
	if err != nil {
		db.logger.Errorw("Error while creating admin", zap.Error(err))
		return err
	}

	db.logger.Infow("admin created", zap.String("name", admin.Name))
	return nil
}

func (db *DB) DeleteAdmin(id string) error {
	admin, err := db.GetAdmin(id)
	if err != nil {
		db.logger.Error("Error retrieving admin", zap.String("id", id), zap.Error(err))
		return err
	}

	db.db.MustExec(deleteAdminQuery, admin.Name)

	return nil
}
