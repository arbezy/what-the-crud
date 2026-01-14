package main

import (
	"github.com/arbezy/what-the-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/reviews", listReviewsHandler)
	router.POST("/reviews", createReviewsHandler)
	router.GET("/reviews/:id", createReviewsHandler)
	router.Run("localhost:8085")
}

func listReviewsHandler(c *gin.Context) {
	reviews := models.ListReviewsHandler()
	if reviews == nil || len(reviews) == 0 {
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
