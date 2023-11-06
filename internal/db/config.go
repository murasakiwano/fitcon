package db

import "fmt"

const (
	DbName          = "fitcon.db"
	FitConnersTable = "fitconners"
)

var (
	fitconnerSchema = fmt.Sprintf(`
create table if not exists %s (
	id varchar(7) primary key,
	name text not null,
	hashed_password text default '',
	team_name text not null,
	team_number integer not null,
	goal1_fat_percentage text default '-',
	goal1_lean_mass text default '-',
	goal1_weight text default '-',
	goal2_fat_percentage text default '-',
	goal2_lean_mass text default '-',
	goal2_visceral_fat text default '-',
	goal2_weight text default '-'
);`, FitConnersTable)
	fitconnerDrop        = fmt.Sprintf("drop table %s;", FitConnersTable)
	getQuery             = fmt.Sprintf("SELECT * FROM %s WHERE id=$1", FitConnersTable)
	insertFitConnerQuery = fmt.Sprintf(`
INSERT INTO %s (
	id,
	name,
	hashed_password,
	team_name,
	team_number,
	goal1_fat_percentage,
	goal1_lean_mass,
	goal1_weight,
	goal2_fat_percentage,
	goal2_lean_mass,
	goal2_visceral_fat,
	goal2_weight
)
VALUES (
	:id,
	:name,
	:hashed_password,
	:team_name,
	:team_number,
	:goal1_fat_percentage,
	:goal1_lean_mass,
	:goal1_weight,
	:goal2_fat_percentage,
	:goal2_lean_mass,
	:goal2_visceral_fat,
	:goal2_weight
);`, FitConnersTable)
)

var (
	INVALID_LENGTH          = &ValidationError{Message: "id must have 7 characters", Code: 0}
	INVALID_FIRST_CHARACTER = &ValidationError{Message: "id must start with 'C' or 'c'", Code: 1}
	PARSE_ERROR             = &ValidationError{Message: "could not convert %v into a number", Code: 2}
)

var updateFitConnerQuery string = `
UPDATE %s
SET %s = ?
WHERE id = ?;`
