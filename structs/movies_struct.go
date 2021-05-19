package structs

// Movies struct which contains
// an array of movies
type Movies struct {
	Movies []Movie `json:"movies"`
}

// Movie struct which contains id, title,
// poster, description, year and actors
type Movie struct {
	Id          uint16  `json:"id"`
	Title       string  `json:"title"`
	Poster      string  `json:"poster"`
	Description string  `json:"description"`
	Year        uint16  `json:"year"`
	Actors      []Actor `json:"actors"`
}

func (m Movie) Print_movie() {

}

// Actor struct which contains a name
type Actor struct {
	Name string `json:"name"`
}
