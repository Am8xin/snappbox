package models

import "time"

type Point struct {
	Latitude   float64
	Longitude  float64
	Timestamp  int64
	DeliveryId int
}

func (p Point) GetTimeStamp() time.Time {
	return time.Unix(0, p.Timestamp*int64(time.Second))
}