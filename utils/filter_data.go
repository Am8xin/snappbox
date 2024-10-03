package utils

import (
	"snappbox_challenge/models"
	"sync"
)

func filterEachList(listOfPoints []models.Point, id int, validCollection map[int][]models.Point, mu *sync.Mutex) {
	
	for j := 1; j < len(listOfPoints); j++ {
		prev := listOfPoints[j-1]
		current := listOfPoints[j]
		distance := CalculateHaversineDistance(&prev, &current)

		timeDiff := current.GetTimeStamp().Sub(prev.GetTimeStamp()).Hours()

		speed := CalculateSpeed(distance, timeDiff)

		if speed <= 100 {
			// CalculateFare(speed, prev.GetTimeStamp(), current.GetTimeStamp()) // we can do this here but we won't because of clean code
			mu.Lock()
			validCollection[id] = append(validCollection[id], current)
			mu.Unlock()
		}
	}
}

func FilterData(pointsCollection *map[int][]models.Point) map[int][]models.Point {
	validCollection := make(map[int][]models.Point)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for id, listOfPoints := range *pointsCollection {
		wg.Add(1)
		validCollection[id] = append(validCollection[id], listOfPoints[0])

		go func(listOfPoints []models.Point, id int) {
			defer wg.Done()
			filterEachList(listOfPoints, id, validCollection, &mu)
		}(listOfPoints, id)
	}

	wg.Wait()
	return validCollection
}
