package dto

type Team struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type TeamCreate struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
