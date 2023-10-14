package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/murasakiwano/fitcon/fitconner"
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

	if len(row) < 9 {
		for i := 9 - len(row); i <= 9; i++ {
			row = append(row, "")
		}
	}

	return fitconner.FitConner{
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

func populateDB(fcs []fitconner.FitConner) error {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	for _, fitConner := range fcs {
		sugar.Infow("creating fitConner...", zap.Object("fitConner", fitConner))
		resp, err := http.PostForm("http://localhost:1323/users", url.Values{
			"teamName":           {fitConner.TeamName},
			"name":               {fitConner.Name},
			"matricula":          {fitConner.ID},
			"goal1FatPercentage": {fitConner.Goal1FatPercentage},
			"goal1LeanMass":      {fitConner.Goal1LeanMass},
			"goal2FatPercentage": {fitConner.Goal2FatPercentage},
			"goal2LeanMass":      {fitConner.Goal2LeanMass},
			"goal2VisceralFat":   {fitConner.Goal2VisceralFat},
		})
		if err != nil {
			sugar.Errorw("error creating fitConner", zap.Error(err))
			return err
		}

		sugar.Infow("successfully created fitConner", zap.Object("fitConner", fitConner), zap.Any("response", resp))
	}

	return nil
}
