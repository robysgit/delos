package requests

type UpdatePondRequest struct {
	ID       string `json:"id"`
	PondName string `json:"pond_name"`
	FarmId   string `json:"farm_id"`
}
