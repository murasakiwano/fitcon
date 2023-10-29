package db

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/murasakiwano/fitcon/internal/auth"
	"github.com/murasakiwano/fitcon/internal/fitconner"
	"github.com/stretchr/testify/assert"
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

var adminTestSchema = Schema{
	create: adminSchema,
	drop:   dropAdmins,
}

var logger, _ = zap.NewDevelopment()

func TestInsertFitConner(t *testing.T) {
	defaultSchema.Create()
	defer defaultSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}
	fc, _ := fitconner.New(
		"C123456",
		"John Doe",
		"test-password1",
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

	hashedPassword, err := auth.HashPassword("password-123")
	if err != nil {
		t.Fatalf("Error hashing password: %s", err)
	}
	db := DB{db: sldb, logger: logger.Sugar()}
	db.db.MustExec(insertFitConnerQuery, "C123456", "John Doe", hashedPassword, "Team 1", 1, "10", "20", "30", "40", "50")

	_, err = db.GetFitConner("C123456")
	if err != nil {
		t.Errorf("Error while getting fitconner: %v", err)
	}
}

func TestBatchInsertFitConner(t *testing.T) {
	defaultSchema.Create()
	defer defaultSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}

	zecaP, _ := fitconner.New("C234142", "Zeca Pagodinho", "test-password1", "Unidos da Brahma", "14", "42", "14", "18", "14", 1)
	monkeyD, _ := fitconner.New("C234143", "Monkey D. Luffy", "test-password2", "Mugiwara", "4", "2", "1", "8", "1", 1)
	roroZ, _ := fitconner.New("C234144", "Roronoa Zoro", "test-password3", "Mugiwara", "2", "1", "-2", "-129", "-2", 1)
	fcs := []fitconner.FitConner{*zecaP, *monkeyD, *roroZ}

	err := db.BatchInsert(fcs)
	if err != nil {
		t.Errorf("Error while batch inserting fitconners: %v", err)
	}

	newFcs := []fitconner.FitConner{}
	err = db.db.Select(&newFcs, "SELECT * FROM fitconners")
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

func TestUpdateFitConner(t *testing.T) {
	defaultSchema.Create()
	defer defaultSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}
	fc, _ := fitconner.New(
		"C123456",
		"John Doe",
		"test-password1",
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

	fields := map[string]string{
		"goal2_lean_mass": "Aumentar 2kg",
	}

	updatedFc, err := db.UpdateFitConner("C123456", fields)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Aumentar 2kg", updatedFc.Goal2LeanMass)
}

func TestInsertAdmin(t *testing.T) {
	adminTestSchema.Create()
	defer adminTestSchema.Drop()
	db := DB{db: sldb, logger: logger.Sugar()}
	hashedPassword, err := auth.HashPassword("password-123")
	if err != nil {
		t.Fatalf("Error hashing password: %s", err)
	}
	admin := Admin{
		Name:           "Zezao",
		HashedPassword: hashedPassword,
	}
	err = db.CreateAdmin(admin)
	if err != nil {
		t.Errorf("Error while creating admin: %v", err)
	}
}

func TestGetAdmin(t *testing.T) {
	adminTestSchema.Create()
	defer adminTestSchema.Drop()

	hashedPassword, err := auth.HashPassword("password-123")
	if err != nil {
		t.Fatalf("Error hashing password: %s", err)
	}
	db := DB{db: sldb, logger: logger.Sugar()}
	db.db.MustExec(insertAdminQuery, "Zezao", hashedPassword)

	admin, err := db.GetAdmin("Zezao")
	if err != nil {
		t.Errorf("Error while getting admin: %v", err)
	}

	assert.Equal(t, admin.Name, "Zezao")
}
