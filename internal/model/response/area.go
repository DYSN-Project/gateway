package response

type Area struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
	Date
}

type AreaList []*Area

func NewAreaIngot() *Area {
	return &Area{}
}
