package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/th1enq/location_recommender/internal/db"
)

func SetupRouter(db *db.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"user":   "th1enq",
		})
	})

	return router
}
