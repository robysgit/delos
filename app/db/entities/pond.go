package entities

import "gorm.io/gorm"

type PondEntity struct {
	gorm.Model
	ID        string          `gorm:"primaryKey"`
	Name      string          `gorm:"uniqueIndex"`
	CreatedBy string          `gorm:"not null"`
	UpdatedBy string          `gorm:"not null"`
	IsDeleted bool            `gorm:"default:false"`
	FarmVPond FarmVPondEntity `gorm:"foreignKey:PondID;references:ID"`
}
