package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// h_input "github.com/callapa1/movies_api/helpers/input"
	// main "github.com/callapa1/movies_api"
	h_print "github.com/callapa1/movies_api/helpers/print"
	structs "github.com/callapa1/movies_api/structs"
)

func Ready() []byte {
	// Open jsonFile, handle error and read
	jsonFile, err := os.Open("movies.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("(Successfully opened movies.json)")
		fmt.Println()
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func Fetch_all() {
	byteValue := Ready()
	// We unmarshal our byteArray which contains our
	// jsonFile's content into 'movies' which we define in /structs
	var movies structs.Movies
	json.Unmarshal(byteValue, &movies)

	// We iterate through every movie within our movies
	// array and print out their details
	for i := 0; i < len(movies.Movies); i++ {

		h_print.Print_id(movies.Movies[i].Id)
		h_print.Print_title(movies.Movies[i].Title)
		h_print.Print_poster(movies.Movies[i].Poster)
		h_print.Print_description(movies.Movies[i].Description)
		h_print.Print_year(movies.Movies[i].Year)

		if len(movies.Movies[i].Actors) > 0 {
			fmt.Println("-Actors: ")
			for _, actor := range movies.Movies[i].Actors {
				h_print.Print_actor(actor.Name)
			}
		} else {
			fmt.Println("-(Actors info not found)")
		}
	}
}
