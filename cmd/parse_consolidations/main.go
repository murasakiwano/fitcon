package main

import (
	"fmt"
	"strconv"

	"github.com/murasakiwano/fitcon/db"
	"github.com/murasakiwano/fitcon/fitconner"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

func main() {
	f, err := excelize.OpenFile("Consolidacoes.xlsx")
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

	fcs := make([]fitconner.Fitconner, 0)
	for idx, row := range rows[3:] {
		if row[0] == "" {
			row[0] = rows[idx+2][0]
		}
		if row[1] == "" {
			row[1] = rows[idx+2][1]
		}
		fitconner := unpackIntoFitconner(row)
		fcs = append(fcs, fitconner)
	}

	if err := populateDB(fcs); err != nil {
		fmt.Printf("error: %+#v\n", err)
	}
}

func unpackIntoFitconner(row []string) fitconner.Fitconner {
	teamNumber, err := strconv.Atoi(row[0])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if len(row) < 9 {
		for i := 9 - len(row); i <= 9; i++ {
			row = append(row, "")
		}
	}

	return fitconner.Fitconner{
		TeamNumber:         teamNumber,
		TeamName:           row[1],
		ID:                 row[2],
		Name:               row[3],
		Goal1FatPercentage: row[4],
		Goal1LeanMass:      row[5],
		Goal2VisceralFat:   row[6],
		Goal2FatPercentage: row[7],
		Goal2LeanMass:      row[8],
	}
}

func populateDB(fcs []fitconner.Fitconner) error {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	fitconners, err := db.New(sugar, "../../fitcon.db")
	fitconners.Drop()
	fitconners.Create()
	defer fitconners.CloseDB()
	if err != nil {
		sugar.Errorw("error", zap.Error(err))
		return err
	}

	if err := fitconners.BatchInsert(fcs); err != nil {
		sugar.Errorw("error", zap.Error(err))
		return err
	}

	return nil
}
