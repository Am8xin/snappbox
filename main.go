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
	
	utils.ReadData("./sample_test.csv", &pointsCollection)
	
	valids := utils.FilterData(&pointsCollection)
	
	totalFares := utils.CalculateTotalFare(valids)
	
	utils.WriteToCSV(totalFares)

	// for measuring the running time of program 
	var t2 time.Time = time.Now()
	
	fmt.Println("Time duration of the program : ", t2.Sub(t).Seconds())
}