package model

type BookRequest struct {
	Name         string `json:"name"`
	Author       string ` json:"author"`
	Release_year int    `json:"release_year"`
	ISBN         string `json:"isbn"`
}
