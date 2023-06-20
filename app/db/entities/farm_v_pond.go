package entities

import "gorm.io/gorm"

type FarmVPondEntity struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FarmID    string `gorm:"index:,unique,composite:farm_pond_uq"`
	PondID    string `gorm:"index:,unique,composite:farm_pond_uq"`
	IsDeleted bool   `gorm:"default:false"`
}
