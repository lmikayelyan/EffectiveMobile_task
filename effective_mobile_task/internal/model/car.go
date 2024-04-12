package model

type Car struct {
	ID      uint
	RegNum  string
	Mark    string
	Model   string
	Year    int
	OwnerID uint
}

type CarInput struct {
	RegNums []string `json:"regNums"`
}

type CarUpdate struct {
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	OwnerID int    `json:"ownerID"`
}
