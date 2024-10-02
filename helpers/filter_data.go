package helpers

import (
	"fmt"

	"snappbox_challenge/models"
)

func FilterData(points []models.Point) map[int][]models.Point {
	fmt.Printf("Count: %v\n", len(points))
	collections := make(map[int][]models.Point)
	validCollection := make(map[int][]models.Point)

	for _, point := range points {
		collections[point.DeliveryId] = append(collections[point.DeliveryId], point)
	}

	for id, listOfPoints := range collections {

		validCollection[id] = append(validCollection[id], listOfPoints[0])

		for j := 1; j < len(listOfPoints); j++ {
			prev := listOfPoints[j-1]
			current := listOfPoints[j]
			distance := CalculateHaversineDistance(&prev, &current)

			timeDiff := current.GetTimeStamp().Sub(prev.GetTimeStamp()).Hours()

			speed := CalculateSpeed(distance, timeDiff)

			// fmt.Printf("Time: %v\n", current.GetTimeStamp())

			// fmt.Printf("Speed: %.2f id: %d and timestamp: %d\n", speed, current.DeliveryId, current.Timestamp)

			if speed <= 100 {
				CalculateFare(speed, prev.GetTimeStamp(), current.GetTimeStamp())
				validCollection[id] = append(validCollection[id], current)
			}

		}

	}

	return validCollection

}