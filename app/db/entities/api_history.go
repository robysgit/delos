package entities

import "gorm.io/gorm"

type ApiHistoryEntity struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Method    string `gorm:"not null"`
	Url       string `gorm:"index"`
	UserAgent string `gorm:"not null"`
}

type ApiHistorySummary struct {
	Method     string
	Url        string
	Total      int
	UniqueUser int
}
