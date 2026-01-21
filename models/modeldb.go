package models

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func getConfig() *mysql.Config {
	var cfg = mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "moviereviews"

	return cfg
}

func ListReviewsHandler() []MovieReview {
	db, err := sql.Open("mysql", getConfig().FormatDSN())
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM moviereviews") // remove this hardcoded value
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	reviews := []MovieReview{}
	for results.Next() {
		var rev MovieReview
		var date string

		err = results.Scan(&rev.ID, &rev.MovieName, &rev.Rating, &date)
		if err != nil {
			panic(err.Error())
		}

		layout := "2006-01-02 15:04:05.000000"
		t, _ := time.Parse(layout, date)
		rev.Date = t

		reviews = append(reviews, rev)
	}

	return reviews
}

func CreateReviewHandler(review MovieReview) {
	db, err := sql.Open("mysql", getConfig().FormatDSN())
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO moviereviews (id, moviename, rating, date) VALUES (?, ?, ?, ?)",
		review.ID, review.MovieName, review.Rating, review.Date,
	)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	insert.Close()
}

func GetReviewByID(id string) *MovieReview {
	db, err := sql.Open("mysql", getConfig().FormatDSN())
	rev := &MovieReview{}
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM moviereviews WHERE id=?", id)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&rev.ID, &rev.MovieName, &rev.Rating, &rev.Date)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return rev
}

func UpdateReviewRating(id string, rating int) *MovieReview {
	if rating < 1 || rating > 10 {
		fmt.Println("Rating is not between 1 and 10")
		return nil
	}

	db, err := sql.Open("mysql", getConfig().FormatDSN())
	rev := &MovieReview{}
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("UPDATE moviereviews SET rating=? WHERE id=?", rating, id)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&rev.ID, &rev.MovieName, &rev.Rating, &rev.Date)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return rev
}
