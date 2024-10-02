package main

import (
	"snappbox_challenge/helpers"
	"snappbox_challenge/utils"
)

func main() {

	points := utils.DataReader("./sample_data.csv")

	valids := helpers.FilterData(points)

	totalFares := helpers.CalculateTotalFare(valids)

	helpers.WriteToCSV(totalFares)

}