package tests

import (
	"net/http"
	"net/http/httptest"
	"serv-e/internal"
	"strings"
	"testing"
)

func TestCreateRecord(t *testing.T) {
	var ds internal.DataStore

	req, err := http.NewRequest("GET", "/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server := http.HandlerFunc(internal.CreateRecordHandler(&ds))
	server.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returnd wrong status code: get %v want %v", status, http.StatusCreated)
	}

	if rr.Body.String() != internal.OKResponseBodyMessage {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), internal.OKResponseBodyMessage)
	}

	if len(ds.GetRecords()) != 1 {
		t.Errorf("data store contains the wrong number of records: got %v want %v", len(ds.GetRecords()), 1)
	}
}

func TestGetRecords(t *testing.T) {
	var ds internal.DataStore
	ds.InsertRecord(internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"})

	req, err := http.NewRequest("GET", "/records", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server := http.HandlerFunc(internal.GetRecordsHandler(&ds))
	server.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: get %v want %v", status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), ds.GetRecords()[0].Id) {
		t.Errorf("handler didn't returned the created record")
	}
}
