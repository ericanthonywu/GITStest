package Model

type AddAuthorsRequest struct {
	BookId     uint   `json:"book_id"`
	AuthorName string `json:"author_name"`
}

type UpdateAuthorsRequest struct {
	AddAuthorsRequest
	DeleteAuthorsRequest
}

type DeleteAuthorsRequest struct {
	Id uint `json:"id"`
}
