package Model

type AddBookRequest struct {
	BookName    string `json:"book_name"`
	PublisherId uint   `json:"publisher_id"`
}

type UpdateBookRequest struct {
	DeleteBookRequest
	AddBookRequest
}

type DeleteBookRequest struct {
	Id uint `json:"id"`
}
