package models

type Todo struct {
	ID          int64  `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
