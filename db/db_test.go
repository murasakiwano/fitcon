package db

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/murasakiwano/fitcon/fitconner"
	"go.uber.org/zap"
)

var sldb *sqlx.DB

func init() {
	var err error
	sldb, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("Error connecting to the database", err)
	}
}

type Schema struct {
	create string
	drop   string
}

func (s Schema) Create() {
	sldb.MustExec(s.create)
}

func (s Schema) Drop() {
	sldb.MustExec(s.drop)
}

var defaultSchema = Schema{
	create: fitconnerSchema,
	drop:   fitconnerDrop,
}

var logger, _ = zap.NewDevelopment()

func TestInsertFitConner(t *testing.T) {
	defaultSchema.Create()
	defer defaultSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}
	fc := fitconner.New(
		"C123456",
		"John Doe",
		"Team 1",
		"10",
		"20",
		"30",
		"40",
		"50",
		1,
	)
	err := db.CreateFitConner(*fc)
	if err != nil {
		t.Errorf("Error while creating fitconner: %v", err)
	}
}

func TestGetFitConner(t *testing.T) {
	defaultSchema.Create()
	defer defaultSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}
	db.db.MustExec(insertFitconnerQuery, "C123456", "John Doe", "Team 1", 1, "10", "20", "30", "40", "50")

	_, err := db.GetFitConner("C123456")
	if err != nil {
		t.Errorf("Error while getting fitconner: %v", err)
	}
}

func TestBatchInsertFitConner(t *testing.T) {
	defaultSchema.Create()
	defer defaultSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}

	fcs := []fitconner.Fitconner{
		*fitconner.New("C234142", "Zeca Pagodinho", "Unidos da Brahma", "14", "42", "14", "18", "14", 1),
		*fitconner.New("C234143", "Monkey D. Luffy", "Mugiwara", "4", "2", "1", "8", "1", 1),
		*fitconner.New("C234144", "Roronoa Zoro", "Mugiwara", "2", "1", "-2", "-129", "-2", 1),
	}

	err := db.BatchInsert(fcs)
	if err != nil {
		t.Errorf("Error while batch inserting fitconners: %v", err)
	}

	newFcs := []fitconner.Fitconner{}
	err = db.db.Select(&newFcs, "SELECT * FROM fitcon_metas")
	if err != nil {
		t.Errorf("Error while getting fitconners: %v", err)
	}

	sort.Slice(fcs, func(i, j int) bool {
		return i > j
	})
	sort.Slice(newFcs, func(i, j int) bool {
		return i > j
	})

	if !reflect.DeepEqual(fcs, newFcs) {
		t.Fatalf("Expected %v, got %v", fcs, newFcs)
	}
}
