package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/teris-io/shortid"
)

var store = make(map[string]string)

type long struct {
	Long_url string `json:"long_url"`
}

var c = false

func Urlshortner(w http.ResponseWriter, r *http.Request) {

	var url long
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		log.Fatalf("failed to decode r.Body into long")
	}
	// logic for same code for same url
	for key := range store {
		if url.Long_url == store[key] {
			c = true
			log.Println("Generated Code : ", key)
		}
	}

	if !c {
		// The package `shortid` enables the generation of short,
		// fully unique, non-sequential and by default URL friendly Ids
		// at a rate of hundreds of thousand per second. It guarantees uniqueness
		//during the time period until 2050!
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			log.Fatalln("Error while generating code")
		}
		urlCode, idErr := sid.Generate()
		if idErr != nil {
			log.Fatalf("error while generating unique number")
		}
		log.Println("Generated Code : ", urlCode)
		store[urlCode] = url.Long_url //storing the original url corresponding to the generated ID
	}

}

func Redirecturl(w http.ResponseWriter, r *http.Request) {

	var actual_url string
	path := r.URL.Path
	p := strings.Split(path, "/")
	actual_url = store[p[1]]
	log.Println("actual_url : ", actual_url)
	http.Redirect(w, r, actual_url, http.StatusPermanentRedirect) // redirect to original url corresponding to the generated ID

}
