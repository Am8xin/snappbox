package utils

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

		if speed <= 100 {
			(*validCollection)[id] = append((*validCollection)[id], current)
		}
	}
	doneChan <- true

}

func FilterData(pointsCollection *map[int][]models.Point) map[int][]models.Point {
	validCollection := make(map[int][]models.Point)

	var i int = 0

	doneChans := make([]chan bool, len(*pointsCollection))

	for id, listOfPoints := range *pointsCollection {

		doneChans[i] = make(chan bool)

		validCollection[id] = append(validCollection[id], listOfPoints[0])

		go filterEachList(listOfPoints, id, &validCollection, doneChans[i])

		<-doneChans[i]
		close(doneChans[i])

		i++
	}

	return validCollection

}