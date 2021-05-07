package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCidHandler(t *testing.T) {
	// req, err := http.NewRequest("get", "/cidNext/?cid=UC_gUM8rL-Lrg6O3adPW9K1g&", nil)
	req, err := http.NewRequest("get", "/cidNext/?cid=UC_gUM8rL-Lrg6O3adPW9K1g&p=CCAQAA", nil)
	// req, err := http.NewRequest("get", "/cidNext/?cid=UC_gUM8rL-Lrg6O3adPW9K1g&p=CBAQAA", nil)
	// req, err := http.NewRequest("get", "/cid/UC_gUM8rL-Lrg6O3adPW9K1g", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	h := http.HandlerFunc(makeHandler(cidNextHandler))
	// h := http.HandlerFunc(makeHandler(cidHandler))
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
