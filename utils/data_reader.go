package utils

import (
	"encoding/csv"
	"log"
	"os"
	"snappbox_challenge/models"
	"strconv"
	"strings"
)

func ReadData(csvFilePath string, pointsCollection *map[int][]models.Point) {
	f, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal("Unable to read input file:", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	// reading file's header
	if _, err := csvReader.Read(); err != nil {
		log.Fatal("Unable to read header:", err)
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("Unable to read record:", err)
		}

		if len(record) < 4 {
			log.Println("Skipping record with insufficient fields:", record)
			continue
		}

		deliveryId, err := strconv.Atoi(strings.TrimSpace(record[0]))
		if err != nil {
			log.Println("Unable to read deliveryId:", err)
			continue
		}

		latitude, err := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		if err != nil {
			log.Println("Unable to read latitude:", err)
			continue
		}

		longitude, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
		if err != nil {
			log.Println("Unable to read longitude:", err)
			continue
		}

		timestamp, err := strconv.ParseInt(strings.TrimSpace(record[3]), 10, 64)
		if err != nil {
			log.Println("Unable to read timestamp:", err)
			continue
		}

		point := models.Point{
			DeliveryId: deliveryId,
			Latitude:   latitude,
			Longitude:  longitude,
			Timestamp:  timestamp,
		}

		(*pointsCollection)[point.DeliveryId] = append((*pointsCollection)[point.DeliveryId], point)
	}
}
