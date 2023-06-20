package ponds

import (
	db "delos/app/db"
	entities "delos/app/db/entities"

	uuid "github.com/google/uuid"
)

func GetPondById(id *string) *entities.PondEntity {
	var pond entities.PondEntity
	if err := db.GetDB().Preload("FarmVPond").Where("ID = ? AND is_deleted = ?", *id, false).Find(&pond).Error; err != nil {
		return nil
	}
	return &pond
}

func GetPonds() []entities.PondEntity {
	var ponds []entities.PondEntity
	db.GetDB().Preload("FarmVPond").Where("is_deleted = ?", false).Find(&ponds)
	return ponds
}

func CreateNewPond(pond *entities.PondEntity, farm_id *string) *entities.PondEntity {
	pond.ID = uuid.New().String()
	db := db.GetDB()
	tx := db.Begin()
	if err := db.Create(pond).Error; err != nil {
		tx.Rollback()
		return nil
	}
	if err := db.Create(&entities.FarmVPondEntity{ID: uuid.NewString(), FarmID: *farm_id, PondID: pond.ID}).Error; err != nil {
		tx.Rollback()
		return nil
	}
	tx.Commit()
	return pond
}

func DeletePondById(id *string) *entities.PondEntity {
	pond := entities.PondEntity{ID: *id, IsDeleted: false}
	db := db.GetDB()
	tx := db.Begin()
	result := db.Model(&pond).Updates(&entities.PondEntity{IsDeleted: true})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil
	}
	result2 := db.Model(&entities.FarmVPondEntity{PondID: *id}).Updates(&entities.FarmVPondEntity{IsDeleted: true})
	if result2.RowsAffected == 0 {
		tx.Rollback()
		return nil
	}
	tx.Commit()
	return &pond
}

func UpdatePond(pond *entities.PondEntity) *entities.PondEntity {
	result := db.GetDB().Model(&pond).Where("is_deleted = ?", false).Updates(&entities.PondEntity{Name: pond.Name})
	if result.RowsAffected == 0 {
		return nil
	}
	return pond
}
