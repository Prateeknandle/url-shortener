package handlers

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Urlshortner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hashed := sha1.New()
		var longurl string
		json.NewDecoder(r.Body).Decode(&longurl)
		// b, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	log.Fatal("Exit")
		// }
		_, err := io.Copy(
			hashed,
			strings.NewReader(longurl),
		)
		if err != nil {
			log.Fatal("Exit")
		}

		encoded := bytes.NewBuffer([]byte{})
		_, err = io.Copy(
			base64.NewEncoder(base64.URLEncoding, encoded),
			bytes.NewReader(hashed.Sum(nil)),
		)
		if err != nil {
			log.Fatal("Exit")
			//return "", err
		}

		encoded.Truncate(8)
		fmt.Printf(encoded.String())
	}
}
