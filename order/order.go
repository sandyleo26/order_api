package order

type Location []string

//Order is the main model
type CreateRequest struct {
	Origin      Location
	Destination Location
}

type Status string

type CreateResponse struct {
	ID       int
	Distance float64
	Status
}

type TakeRequest struct {
	Status
}

type TakeResponse struct {
	Status
}

type GetOptions struct {
	Page  int
	Limit int
}

type GetResponse CreateResponse

type ErrorResponse struct {
	Error string
}
