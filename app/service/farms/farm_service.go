package farms

import (
	entities "delos/app/db/entities"
	repo "delos/app/db/repositories/farms"
	model "delos/app/service/model"
	"fmt"
)

func GetFarm(id *string) *model.Farm {
	farm := repo.GetFarmById(id)
	if farm == nil {
		return nil
	}
	return &model.Farm{ID: farm.ID, FarmName: farm.Name}
}

func GetFarms() []model.Farm {
	farms := repo.GetFarms()
	var result []model.Farm
	for _, r := range farms {
		result = append(result, model.Farm{ID: r.ID, FarmName: r.Name})
	}
	return result
}

func CreateFarm(farm *model.Farm) *model.Farm {
	entity := repo.CreateNewFarm(&entities.FarmEntity{Name: farm.FarmName, IsDeleted: false})
	if entity == nil {
		return nil
	}
	farm.ID = entity.ID
	return farm
}

func DeleteFarm(id *string) int8 {
	farm := repo.DeleteFarmById(id)
	if farm == nil {
		fmt.Println("Not Found")
		return 0
	}
	return 1
}

func UpdateFarm(farm *model.Farm) *model.Farm {
	eFarm := entities.FarmEntity{ID: farm.ID, Name: farm.FarmName}
	entity := repo.UpdateFarm(&eFarm)
	if entity == nil {
		entity = repo.CreateNewFarm(&eFarm)
	}
	return &model.Farm{ID: entity.ID, FarmName: entity.Name}
}
