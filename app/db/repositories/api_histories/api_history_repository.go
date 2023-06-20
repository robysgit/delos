package apihistories

import (
	db "delos/app/db"
	entities "delos/app/db/entities"

	uuid "github.com/google/uuid"
)

func CreateApiHistory(history *entities.ApiHistoryEntity) *entities.ApiHistoryEntity {
	history.ID = uuid.New().String()
	if err := db.GetDB().Create(history).Error; err != nil {
		return nil
	}
	return history
}

func GetApiHistorySummaries() []entities.ApiHistorySummary {
	var result []entities.ApiHistorySummary
	db.GetDB().Raw("SELECT method, url, COUNT(1) as total, COUNT(DISTINCT user_agent) as unique_user FROM api_history_entities GROUP BY method, url").Scan(&result)
	return result
}
