package models

import (
	"gorm.io/gorm"
	"time"
)

type EmailLog struct {
	gorm.Model
	To          string    `json:"to"`
	From        string    `json:"from"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	Status      string    `json:"status"`  // success, failed, bounced
	EmailType   string    `json:"email_type"`  // transactional, marketing
	SentAt      time.Time `json:"sent_at"`
	RetryCount  int       `json:"retry_count"`
	Latency     float64   `json:"latency"` // Time taken in milliseconds
}
