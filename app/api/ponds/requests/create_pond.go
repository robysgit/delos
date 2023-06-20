package requests

type CreatePondRequest struct {
	PondName string `json:"pond_name"`
	FarmId   string `json:"farm_id"`
}
