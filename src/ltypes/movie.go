package ltypes

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}