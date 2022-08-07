package form

type AreaList struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type AreaShow struct {
	Id string `json:"id"`
}
