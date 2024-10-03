package main

import (
	"fmt"
	"snappbox_challenge/models"
	"snappbox_challenge/utils"
	"time"
)

func main() {

	var t time.Time = time.Now()

	pointsCollection := make(map[int][]models.Point)

	utils.ReadData("./sample_data.csv", &pointsCollection)

/*
	valids := utils.FilterData(&pointsCollection)

	totalFares := utils.CalculateTotalFare(valids)

	utils.WriteToCSV(totalFares)
*/

	// does the exact thing that multi line comment above does but a little faster
	utils.WriteToCSV(utils.CalculateTotalFare(utils.FilterData(&pointsCollection)))

	var t2 time.Time = time.Now()
	fmt.Println(t2.Sub(t).Seconds())
	
}