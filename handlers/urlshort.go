package handlers

import (
	"encoding/json"

	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/teris-io/shortid"
)

//var store = make(map[string]string)

type long struct {
	Long_url string `json:"long_url"`
}

func Split(r rune) bool {
	return r == ';' || r == '='
}

var urlstr = "http://localhost:3000/"

var valid_url = regexp.MustCompile(`^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(((([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)\.)?)?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?)|localhost)(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)

func Urlshortner(w http.ResponseWriter, r *http.Request) {

	var c = false
	var url long
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		log.Fatalf("failed to decode r.Body into long")
	}
	err = Validation(url.Long_url)
	if err != nil {
		log.Fatalln("given long_url is not a valid url")
	}
	// logic for same short url if  same url is passed
	// for key := range store {
	// 	if url.Long_url == store[key] {
	// 		c = true
	// 		log.Println("Generated Short Link : ", urlstr+key)
	// 	}
	// }
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Println("can't read the file")
	}
	set := strings.FieldsFunc(string(data), Split)
	for k, v := range set {
		if v == url.Long_url {
			c = true
			log.Println("Generated Short Link : ", urlstr+set[k-1])
			w.Write([]byte(urlstr + set[k-1]))
		}
	}

	if !c {
		// cheching for file, if not found, it will generate a new file named "data.txt"
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal(err)
		}
		// The package `shortid` enables the generation of short,
		// fully unique, non-sequential and by default URL friendly Ids
		// at a rate of hundreds of thousand per second. It guarantees uniqueness
		// during the time period until 2050!
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			log.Fatalln("Error while generating code")
		}
		urlCode, idErr := sid.Generate()
		if idErr != nil {
			log.Fatalf("error while generating unique number")
		}
		log.Println("Generated Short Link : ", urlstr+urlCode)
		//store[urlCode] = url.Long_url //storing the original url corresponding to the short url
		// storing the new long_url corresponding to its shortened url
		if _, err = f.WriteString(urlCode + "=" + url.Long_url + ";"); err != nil {
			panic(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(urlstr + urlCode))
	}
}

func Redirecturl(w http.ResponseWriter, r *http.Request) {

	var actual_url string
	path := r.URL.Path
	p := strings.Split(path, "/")
	//actual_url = store[p[1]]
	data, _ := os.ReadFile("data.txt")
	set := strings.FieldsFunc(string(data), Split)
	for k, v := range set {
		if v == p[1] {
			actual_url = set[k+1]
		}
	}
	http.Redirect(w, r, actual_url, http.StatusPermanentRedirect) // redirect to original url corresponding to the generated ID

}

func Validation(urlv string) error {
	valid := valid_url.MatchString(urlv)
	if !valid {
		log.Fatalln("given long_url is not a valid url")
	}
	return nil
}
