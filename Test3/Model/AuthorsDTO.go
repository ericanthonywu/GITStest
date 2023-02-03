package Model

type AuthorsRequest struct {
	BookId     uint   `json:"book_id"`
	AuthorName string `json:"author_name"`
}
