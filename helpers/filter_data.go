package helpers

import (
	"snappbox_challenge/models"
)

func filterEachList(listOfPoints []models.Point, id int, validCollection *map[int][]models.Point, doneChan chan bool) {
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
			(*validCollection)[id] = append((*validCollection)[id], current)
		}
	}
	doneChan <- true

}

func FilterData(points []models.Point) map[int][]models.Point {
	collections := make(map[int][]models.Point)
	validCollection := make(map[int][]models.Point)

	// Grouping the points into each id_delivery group
	for _, point := range points {
		collections[point.DeliveryId] = append(collections[point.DeliveryId], point)
	}

	var i int = 0

	doneChans := make([]chan bool, len(collections))

	for id, listOfPoints := range collections {

		doneChans[i] = make(chan bool)

		validCollection[id] = append(validCollection[id], listOfPoints[0])

		go filterEachList(listOfPoints, id, &validCollection, doneChans[i])

		<-doneChans[i]
		close(doneChans[i])

		i++
	}

	return validCollection

}