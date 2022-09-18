package handlers

import (
	"testing"
)

func Test_Url_Validation(t *testing.T) {
	url := []byte("https://github.com/")
	err := Validation(string(url))
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	url = []byte("htt://github.com/")
	err = Validation(string(url))
	if err != nil {
		t.FailNow()
	}
}
