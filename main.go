package main

import (
	//"fmt"
	"snappbox_challenge/models"
	"snappbox_challenge/utils"
	//"time"
)

func main() {

	// var t time.Time = time.Now()

	pointsCollection := make(map[int][]models.Point)

	utils.ReadData("./large_delivery_dataset.csv", &pointsCollection)


	valids := utils.FilterData(&pointsCollection)

	totalFares := utils.CalculateTotalFare(valids)

	utils.WriteToCSV(totalFares)


	// does the exact same thing that multi line comment above does but a little faster
	//utils.WriteToCSV(utils.CalculateTotalFare(utils.FilterData(&pointsCollection)))


	// for measuring the running time of program 
	// var t2 time.Time = time.Now()
	// fmt.Println(t2.Sub(t).Seconds())
	
}