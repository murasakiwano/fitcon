package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/murasakiwano/fitcon/internal/db"
	"github.com/murasakiwano/fitcon/internal/fitconner"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

func main() {
	f, err := excelize.OpenFile("03 Consolidação Metas.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Metas")
	if err != nil {
		fmt.Println(err)
		return
	}

	fcs := make([]fitconner.FitConner, 0)
	for idx, row := range rows[3:] {
		if row[0] == "" {
			row[0] = rows[idx+2][0]
		}
		if row[1] == "" {
			row[1] = rows[idx+2][1]
		}
		fitconner := unpackIntoFitConner(row)
		fcs = append(fcs, fitconner)
	}

	if err := populateDB(fcs); err != nil {
		fmt.Printf("error: %+#v\n", err)
	}
}

func unpackIntoFitConner(row []string) fitconner.FitConner {
	teamNumber, err := strconv.Atoi(row[0])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if len(row) < 11 {
		for len(row) < 11 {
			row = append(row, "*")
		}
	}

	log.Printf("%#+v", row)
	log.Printf("=============ROWLEN=============\n%d\n", len(row))
	fc := fitconner.FitConner{
		TeamNumber:         teamNumber,
		TeamName:           row[1],
		ID:                 row[2],
		Name:               row[3],
		Goal1FatPercentage: row[4],
		Goal1Weight:        row[5],
		Goal1LeanMass:      row[6],
		Goal2VisceralFat:   row[7],
		Goal2Weight:        row[8],
		Goal2FatPercentage: row[9],
		Goal2LeanMass:      row[10],
	}
	log.Printf("FitConner: %+v", fc)

	return fc
}

func populateDB(fcs []fitconner.FitConner) error {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	os.Setenv("DATABASE_FILE", "./fitcon.db")
	db, err := db.New(sugar)
	if err != nil {
		sugar.Error("Error ocurred creating database", zap.Error(err))
		return err
	}

	sugar.Infow("creating fitConners...")
	db.BatchInsert(fcs)
	sugar.Infow("successfully created fitConners")

	db.GetAllFitConners()

	return nil
}
