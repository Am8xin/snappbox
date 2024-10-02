package helpers

import "snappbox_challenge/models"

func calculateEachFare(listOfValidPoints []models.Point, total *float64, doneChan chan bool) {
	for i := 1; i < len(listOfValidPoints); i++ {

		current := listOfValidPoints[i]
		prev := listOfValidPoints[i-1]
		distance := CalculateHaversineDistance(&prev, &current)
		timeDiff := current.GetTimeStamp().Sub(prev.GetTimeStamp()).Hours()
		speed := CalculateSpeed(distance, timeDiff)

		*total += CalculateFare(speed, prev.GetTimeStamp(), current.GetTimeStamp())

	}

	doneChan <- true
}

func CalculateTotalFare(validPoints map[int][]models.Point) map[int]float64 {
	result := make(map[int]float64)

	doneChans := make([]chan bool, len(validPoints))

	var i int = 0

	for id, listOfValidPoints := range validPoints {
		var total float64 = 1.30

		for i := 1; i < len(listOfValidPoints); i++ {
		doneChans[i] = make(chan bool)

			current := listOfValidPoints[i]
			prev := listOfValidPoints[i-1]
			distance := CalculateHaversineDistance(&prev, &current)
			timeDiff := current.GetTimeStamp().Sub(prev.GetTimeStamp()).Hours()
			speed := CalculateSpeed(distance, timeDiff)
		go calculateEachFare(listOfValidPoints, &total, doneChans[i])

			total += CalculateFare(speed, prev.GetTimeStamp(), current.GetTimeStamp())

		}
		<-doneChans[i]

		if total >= 3.47 {
			result[id] = total
		}

		i++
	}
	return result
}