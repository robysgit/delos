package farms

import (
	db "delos/app/db"
	entities "delos/app/db/entities"

	uuid "github.com/google/uuid"
)

func GetFarmById(id *string) *entities.FarmEntity {
	var farm entities.FarmEntity
	if err := db.GetDB().Where("ID = ? AND is_deleted = ?", *id, false).Find(&farm).Error; err != nil {
		return nil
	}
	return &farm
}

func GetFarms() []entities.FarmEntity {
	var farms []entities.FarmEntity
	db.GetDB().Where("is_deleted = ?", false).Find(&farms)
	return farms
}

func CreateNewFarm(farm *entities.FarmEntity) *entities.FarmEntity {
	farm.ID = uuid.New().String()
	if err := db.GetDB().Create(farm).Error; err != nil {
		return nil
	}
	return farm
}

func DeleteFarmById(id *string) *entities.FarmEntity {
	farm := entities.FarmEntity{ID: *id, IsDeleted: false}
	db := db.GetDB()
	tx := db.Begin()
	result := db.Model(&farm).Updates(&entities.FarmEntity{IsDeleted: true})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil
	}
	result2 := db.Model(&entities.FarmVPondEntity{FarmID: *id}).Updates(&entities.FarmVPondEntity{IsDeleted: true})
	if result2.RowsAffected == 0 {
		tx.Rollback()
		return nil
	}
	tx.Commit()
	return &farm
}

func UpdateFarm(farm *entities.FarmEntity) *entities.FarmEntity {
	result := db.GetDB().Model(&farm).Where("is_deleted = ?", false).Updates(&entities.FarmEntity{Name: farm.Name})
	if result.RowsAffected == 0 {
		return nil
	}
	return farm
}
