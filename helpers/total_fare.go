package helpers

import "snappbox_challenge/models"


func CalculateTotalFare(validPoints map[int][]models.Point) map[int]float64 {
	result := make(map[int]float64)

	for id, listOfValidPoints := range validPoints {
		var total float64 = 1.30

		for i := 1; i < len(listOfValidPoints); i++ {

			current := listOfValidPoints[i]
			prev := listOfValidPoints[i-1]
			distance := CalculateHaversineDistance(&prev, &current)
			timeDiff := current.GetTimeStamp().Sub(prev.GetTimeStamp()).Hours()
			speed := CalculateSpeed(distance, timeDiff)

			total += CalculateFare(speed, prev.GetTimeStamp(), current.GetTimeStamp())

		}

		if total >= 3.47 {
			result[id] = total
		}

	}
	return result
}