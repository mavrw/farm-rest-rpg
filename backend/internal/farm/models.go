package farm

type CreateFarmInput struct {
	Name string `json:"name" binding:"required"`
}
