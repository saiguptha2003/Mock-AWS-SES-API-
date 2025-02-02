package routes

import (
	"mock-ses-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/v1/email")
	{
		api.POST("/send", func(c *gin.Context) { controllers.SendEmail(c, db) })
		api.GET("/statistics", func(c *gin.Context) { controllers.GetStatistics(c, db) })
		api.GET("/search", func(c *gin.Context) { controllers.GetEmails(c, db) })
		api.POST("/retry", func(c *gin.Context) { controllers.RetryFailedEmails(c, db) })
		api.POST("/send-bulk", func(c *gin.Context) { controllers.SendBulkEmails(c, db) })
		api.GET("/stats/sender", func(c *gin.Context) { controllers.GetSenderStatistics(c, db) })
	}
}

