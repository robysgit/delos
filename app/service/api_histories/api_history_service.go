package apihistories

import (
	repo "delos/app/db/repositories/api_histories"
	model "delos/app/service/model"
)

func GetApiHistorySummaries() []model.ApiHistorySummary {
	data := repo.GetApiHistorySummaries()
	var result []model.ApiHistorySummary
	for _, r := range data {
		result = append(result, model.ApiHistorySummary{Method: r.Method, Url: r.Url, Total: r.Total, UniqueUser: r.UniqueUser})
	}
	return result
}
