package Model

type AddPublisherRequest struct {
	PublisherName string `json:"publisher_name"`
}

type UpdatePublisherRequest struct {
	DeleteBookRequest
	AddPublisherRequest
}

type DeletePublisherRequest struct {
	Id uint `json:"id"`
}
