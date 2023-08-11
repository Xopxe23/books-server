package domain

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

type UpdateBookInput struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Rating *int    `json:"rating"`
}
