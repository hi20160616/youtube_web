package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCidHandler(t *testing.T) {
	req, err := http.NewRequest("get", "/cid/UC-5VbWqa7FfpDaK2lLwE3dg", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	h := http.HandlerFunc(makeHandler(cidHandler))
	h.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Error(status)
	}

	fmt.Println(w.Body)
}

func TestVidHandler(t *testing.T) {
	req, err := http.NewRequest("get", "/vid/S1zL6m8Azuw", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	h := http.HandlerFunc(makeHandler(vidHandler))
	h.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Error(status)
	}

	fmt.Println(w.Body)
}
