package impl

import (
	entities "delos/app/db/entities"
	repo "delos/app/db/repositories/ponds"
	model "delos/app/service/model"
)

func GetPond(id *string) *model.Pond {
	pond := repo.GetPondById(id)
	if pond == nil {
		return nil
	}
	return &model.Pond{ID: pond.ID, PondName: pond.Name, FarmId: pond.FarmVPond.FarmID}
}

func GetPonds() []model.Pond {
	ponds := repo.GetPonds()
	var result []model.Pond
	for _, r := range ponds {
		result = append(result, model.Pond{ID: r.ID, PondName: r.Name, FarmId: r.FarmVPond.FarmID})
	}
	return result
}

func CreatePond(pond *model.Pond) *model.Pond {
	entity := repo.CreateNewPond(&entities.PondEntity{Name: pond.PondName, IsDeleted: false}, &pond.FarmId)
	if entity == nil {
		return nil
	}
	pond.ID = entity.ID
	return pond
}

func DeletePond(id *string) int8 {
	pond := repo.DeletePondById(id)
	if pond == nil {
		return 0
	}
	return 1
}

func UpdatePond(pond *model.Pond) *model.Pond {
	ePond := entities.PondEntity{ID: pond.ID, Name: pond.PondName}
	entity := repo.UpdatePond(&ePond)
	if entity == nil {
		entity = repo.CreateNewPond(&ePond, &pond.FarmId)
	}
	return &model.Pond{ID: entity.ID, PondName: entity.Name, FarmId: entity.FarmVPond.FarmID}
}
