package utils

import (
	"snappbox_challenge/models"
	"sync"
)

func calculateEachFare(listOfValidPoints []models.Point, total *float64, mu *sync.Mutex) {
	for i := 1; i < len(listOfValidPoints); i++ {
		current := listOfValidPoints[i]
		prev := listOfValidPoints[i-1]
		distance := CalculateHaversineDistance(&prev, &current)
		timeDiff := current.GetTimeStamp().Sub(prev.GetTimeStamp()).Hours()
		speed := CalculateSpeed(distance, timeDiff)

		fare := CalculateFare(speed, prev.GetTimeStamp(), current.GetTimeStamp())

		mu.Lock()
		*total += fare
		mu.Unlock()
	}
}

func CalculateTotalFare(validPoints map[int][]models.Point) map[int]float64 {
	result := make(map[int]float64)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for id, listOfValidPoints := range validPoints {
		wg.Add(1)
		total := 1.30

		go func(listOfValidPoints []models.Point, id int) {
			defer wg.Done()
			calculateEachFare(listOfValidPoints, &total, &mu)

			if total >= 3.47 {
				mu.Lock()
				result[id] = total
				mu.Unlock()
			}
		}(listOfValidPoints, id)
	}

	wg.Wait()
	return result
}
