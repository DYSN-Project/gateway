package response

type Area struct {
	ID          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Image       string `bson:"image"`
	Status      string `bson:"status"`
	Date
}

type AreaList []*Area

func NewAreaIngot() *Area {
	return &Area{}
}
