package controllers

import (
	"math/rand"
	"mock-ses-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func SendEmail(c *gin.Context, db *gorm.DB) {
	var email models.EmailLog

	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	startTime := time.Now()

	if len(email.To) < 5 {
		email.Status = "bounced"
	} else if time.Now().Unix()%5 == 0 { 
		email.Status = "failed"
	} else {
		email.Status = "success"
	}

	email.SentAt = time.Now()
	email.Latency = float64(time.Since(startTime).Milliseconds())
	db.Create(&email)

	c.JSON(http.StatusOK, gin.H{"message": "Email processed", "email": email})
}

func GetStatistics(c *gin.Context, db *gorm.DB) {
	var totalEmails, successEmails, bouncedEmails, failedEmails int64
	var avgLatency float64
	var topRecipients []string

	db.Model(&models.EmailLog{}).Count(&totalEmails)
	db.Model(&models.EmailLog{}).Where("status = ?", "success").Count(&successEmails)
	db.Model(&models.EmailLog{}).Where("status = ?", "bounced").Count(&bouncedEmails)
	db.Model(&models.EmailLog{}).Where("status = ?", "failed").Count(&failedEmails)
	db.Model(&models.EmailLog{}).Select("AVG(latency)").Scan(&avgLatency)

	db.Raw("SELECT to FROM email_logs GROUP BY to ORDER BY COUNT(*) DESC LIMIT 5").Scan(&topRecipients)

	var failureRate float64
	if totalEmails > 0 {
		failureRate = (float64(bouncedEmails+failedEmails) / float64(totalEmails)) * 100
	}

	c.JSON(http.StatusOK, gin.H{
		"total_emails":    totalEmails,
		"successful":      successEmails,
		"bounced":        bouncedEmails,
		"failed":         failedEmails,
		"average_latency": avgLatency,
		"top_recipients":  topRecipients,
		"failure_rate":    failureRate,
	})
}


func RetryFailedEmails(c *gin.Context, db *gorm.DB) {
	var failedEmails []models.EmailLog
	db.Where("status = ?", "failed").Find(&failedEmails)

	for _, email := range failedEmails {
		if email.RetryCount < 3 {
			email.Status = "success"
			email.RetryCount++
			db.Save(&email)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Retries processed", "emails_retried": len(failedEmails)})
}

func GetEmails(c *gin.Context, db *gorm.DB) {
	var emails []models.EmailLog
	status := c.Query("status")
	to := c.Query("to")
	date := c.Query("date")

	query := db.Model(&models.EmailLog{})

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if to != "" {
		query = query.Where("to = ?", to)
	}
	if date != "" {
		query = query.Where("DATE(sent_at) = ?", date)
	}

	query.Order("sent_at DESC").Find(&emails)

	c.JSON(http.StatusOK, gin.H{"emails": emails})
}


func SendBulkEmails(c *gin.Context, db *gorm.DB) {
	type BulkEmailRequest struct {
		From    string   `json:"from"`
		To      []string `json:"to"`
		Subject string   `json:"subject"`
		Body    string   `json:"body"`
	}

	var request BulkEmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	startTime := time.Now()

	var logs []models.EmailLog
	successCount, failureCount := 0, 0

	for _, recipient := range request.To {
		latency := float64(rand.Intn(300) + 100) 
		status := "success"

		if rand.Float32() < 0.1 {
			status = "failed"
			failureCount++
		} else {
			successCount++
		}

		logs = append(logs, models.EmailLog{
			From:    request.From,
			To:      recipient,
			Subject: request.Subject,
			Body:    request.Body,
			Status:  status,
			Latency: latency,
		})
	}

	db.Create(&logs)
	totalTime := time.Since(startTime).Seconds()

	c.JSON(http.StatusOK, gin.H{
		"message":       "Bulk email processing complete",
		"total_sent":    len(request.To),
		"successful":    successCount,
		"failed":        failureCount,
		"total_latency": totalTime,
	})
}





func GetSenderStatistics(c *gin.Context, db *gorm.DB) {
	sender := c.Query("sender") 

	var totalEmails, successEmails, failedEmails int64
	db.Model(&models.EmailLog{}).Where("from = ?", sender).Count(&totalEmails)
	db.Model(&models.EmailLog{}).Where("from = ? AND status = ?", sender, "success").Count(&successEmails)
	db.Model(&models.EmailLog{}).Where("from = ? AND status = ?", sender, "failed").Count(&failedEmails)

	var avgLatency float64
	db.Model(&models.EmailLog{}).Where("from = ?", sender).Select("AVG(latency)").Scan(&avgLatency)

	c.JSON(http.StatusOK, gin.H{
		"sender":          sender,
		"total_sent":      totalEmails,
		"successful":      successEmails,
		"failed":          failedEmails,
		"average_latency": avgLatency,
	})
}