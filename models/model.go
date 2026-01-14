package models

import (
	"time"
)

type MovieReview struct {
	ID        string    `json:"id"`
	MovieName string    `json:"name"`
	Rating    int       `json:"rating"`
	Date      time.Time `json:"date"`
}

func SampleReviews() []MovieReview {
	var reviews = []MovieReview{
		{ID: "3", MovieName: "Manchester by the Sea", Rating: 10, Date: time.Now()},
	}

	return reviews
}
