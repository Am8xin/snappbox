package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"snappbox_challenge/models"
)

func DataReader(csvFilePath string) []models.Point {

	points := []models.Point{}

	f, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read records ")
	}

	recordsLength := len(records)

	for i := 1; i < recordsLength; i++ {

		deliveryId, err := strconv.ParseInt(records[i][0], 10, 64)
		if err != nil {
			log.Fatal("Unable to read deliveryId")
		}

		latitude, err := strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			log.Fatal("Unable to read latitude")
		}

		longitude, err := strconv.ParseFloat(records[i][2], 64)
		if err != nil {
			log.Fatal("Unable to read longitude")
		}

		timestamp, err := strconv.ParseInt(records[i][3], 10, 64)
		if err != nil {
			log.Fatal("Unable to read timestamp")
		}

		point := models.Point{
			DeliveryId: int(deliveryId),
			Latitude:   latitude,
			Longitude:  longitude,
			Timestamp:  timestamp,
		}

		points = append(points, point)

	}

	return points
}