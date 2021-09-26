package repository

type PaginatorOption struct {
	PerPage     int
	CurrentPage int
	WithTrashed bool
}
type PaginatorPayload struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}
