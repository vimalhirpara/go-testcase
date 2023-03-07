package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {

	expected := "Hello vimal"

	req := httptest.NewRequest(http.MethodGet, "/greet?name=vimal", nil)
	w := httptest.NewRecorder()

	//call actual handler
	RequestHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.Equal(t, expected, string(data))

	// if err != nil {
	// 	t.Errorf("Error: %v", err)
	// }

	// if string(data) != expected {
	// 	t.Errorf("Expected Hello vimal but got %v", string(data))
	// }
}
