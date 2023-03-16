package form

type CreateArea struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      string `json:"status"`
}

type UpdateArea struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      string `json:"status"`
}
