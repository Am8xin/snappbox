package main

import (
	"snappbox_challenge/models"
	"snappbox_challenge/utils"
)

func main() {

	pointsCollection := make(map[int][]models.Point)

	utils.ReadData("./sample_data.csv", &pointsCollection)

	valids := utils.FilterData(&pointsCollection)

	totalFares := utils.CalculateTotalFare(valids)

	utils.WriteToCSV(totalFares)

}