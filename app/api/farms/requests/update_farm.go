package requests

type UpdateFarmRequest struct {
	ID       string `json:"id"`
	FarmName string `json:"farm_name"`
}
