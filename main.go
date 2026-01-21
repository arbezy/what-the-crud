package main

import (
	"net/http"
	"strconv"

	"github.com/arbezy/what-the-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/reviews", listReviewsHandler)
	router.POST("/reviews", createReviewsHandler)
	router.GET("/reviews/:id", getMoviesByID)
	router.PATCH("reviews/:id:rating", updateReviewRating)
	router.Run("localhost:8085")
}

func listReviewsHandler(c *gin.Context) {
	reviews := models.ListReviewsHandler()
	if len(reviews) == 0 { // len on nil returns 0
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, reviews)
	}
}

func createReviewsHandler(c *gin.Context) {
	var rev models.MovieReview

	if err := c.BindJSON(&rev); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.CreateReviewHandler(rev) // NOTE: this should catch an error if this process fails along the way...
		c.IndentedJSON(http.StatusCreated, rev)
	}
}

func getMoviesByID(c *gin.Context) {
	id := c.Param("id")

	rev := models.GetReviewByID(id)
	if rev == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, rev)
	}
}

func updateReviewRating(c *gin.Context) {
	id := c.Param("id")
	ratingStr := c.Param("rating")

	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		c.AbortWithError(500, err)
	}

	rev := models.UpdateReviewRating(id, rating)
	if rev == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, rev)
	}
}
