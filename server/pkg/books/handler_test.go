package books

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooksHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal("Error creating request:", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}

	expected := `{"books":[]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned wrong status code: got %v wanted %v",
			rr.Body.String(), expected)
	}
}
