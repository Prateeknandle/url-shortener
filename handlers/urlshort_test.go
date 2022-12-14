package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Url_Validation(t *testing.T) {
	url := []byte("https://github.com/")
	err := Validation(string(url))
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	url = []byte("http://github.com")
	err = Validation(string(url))
	if err != nil {
		assert.Error(t, err)
	}
}
