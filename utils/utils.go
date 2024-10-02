
package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
	"snappbox_challenge/models"
	"github.com/umahmood/haversine"
)

func CalculateHaversineDistance(p1 *models.Point, p2 *models.Point) float64 {
	p1Coordinates := haversine.Coord{Lat: p1.Latitude, Lon: p1.Longitude}
	p2Coordinates := haversine.Coord{Lat: p2.Latitude, Lon: p2.Longitude}

	_, km := haversine.Distance(p1Coordinates, p2Coordinates)

	return km
}

func CalculateSpeed(distance float64, timeDiff float64) float64 {
	deltaTime := timeDiff
	return distance / deltaTime
}

func inTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) && check.Before(end)) || check.Equal(end)
}

// We calculate the total fare between two points
func CalculateFare(speed float64, prevTimeStamp, currentTimeStamp time.Time) float64 {
	zeroAmTomorrow := time.Date(currentTimeStamp.Year(), currentTimeStamp.Month(), currentTimeStamp.Day()+1, 0, 0, 0, 0, currentTimeStamp.Location())
	zeroAm := time.Date(currentTimeStamp.Year(), currentTimeStamp.Month(), currentTimeStamp.Day(), 0, 0, 0, 0, currentTimeStamp.Location())
	fiveAm := time.Date(currentTimeStamp.Year(), currentTimeStamp.Month(), currentTimeStamp.Day(), 5, 0, 0, 0, currentTimeStamp.Location())
	distance := speed * float64(currentTimeStamp.Sub(prevTimeStamp).Hours())

	if speed <= 10 {
		return distance * 11.90
	}

	var total float64 = 0

	if currentTimeStamp.Sub(prevTimeStamp).Hours() >= 24 {
		daysCount := currentTimeStamp.Sub(prevTimeStamp).Hours() / 24
		total += daysCount * 20.56 // each day delivery cost is a constant value of (5*1.30 + 19*0.74) which is equal to 20.56
		prevTimeStamp = prevTimeStamp.AddDate(0, 0, int(daysCount))
	}

	// if five am is between prev and current
	if inTimeSpan(prevTimeStamp, currentTimeStamp, fiveAm) {
		kmBeforeFiveAm := speed * float64(fiveAm.Sub(prevTimeStamp).Hours())
		kmAfterFiveAm := speed * float64(currentTimeStamp.Sub(fiveAm).Hours())
		total += kmBeforeFiveAm * 1.30
		total += kmAfterFiveAm * 0.74
		return total
	}

	// if zero am is between prev and current
	if inTimeSpan(prevTimeStamp, currentTimeStamp, zeroAm) {
		kmBeforeZeroAm := speed * float64(zeroAm.Sub(prevTimeStamp).Hours())
		kmAfterZeroAm := speed * float64(currentTimeStamp.Sub(zeroAm).Hours())
		total += kmBeforeZeroAm * 0.74
		total += kmAfterZeroAm * 1.30
		return total
	}

	// if prev and current both are between zero am and five am
	if inTimeSpan(zeroAm, fiveAm, currentTimeStamp) && inTimeSpan(zeroAm, fiveAm, prevTimeStamp) {
		return distance * 1.30
	}
	// if prev and current both are between five am and zero am of the next day
	if inTimeSpan(fiveAm, zeroAmTomorrow, currentTimeStamp) && inTimeSpan(fiveAm, zeroAmTomorrow, prevTimeStamp) {
		return distance * 0.74
	}

	return total

}

/*
func WriteToCSV(inputMap map[int]float64) {
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"id_delivery", "fare_estimate"}
	if err := writer.Write(header); err != nil {
		log.Fatalln("Error writing header to CSV:", err)
	}

	for key, value := range inputMap {
		record := []string{strconv.FormatInt(int64(key), 10), strconv.FormatFloat(value, 'f', -1, 64)}
		if err := writer.Write(record); err != nil {
			log.Fatalln("Error writing record to CSV: ", err)
		}
	}
}
*/

func WriteToCSV(inputMap map[int]float64) {  
	file, err := os.Create("output.csv")  
	if err != nil {  
		log.Fatal("Failed to create file:", err)  
	}  
	defer file.Close()  

	writer := csv.NewWriter(file)  
	defer writer.Flush()  

	if err := writer.Write([]string{"id_delivery", "fare_estimate"}); err != nil {  
		log.Fatalln("Error writing header to CSV:", err)  
	}  

	for key, value := range inputMap {  
		if err := writer.Write([]string{  
			strconv.FormatInt(int64(key), 10),  
			strconv.FormatFloat(value, 'f', -1, 64),  
		}); err != nil {  
			log.Fatalln("Error writing record to CSV:", err)  
		}  
	}  
}