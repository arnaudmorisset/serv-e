package internal

import (
	"html/template"
	"io"
	"net/http"
	"time"
)

func CreateRecordHandler(ds *InMemoryDataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
		}

		record := Record{Id: time.Now().Format("15:04:05"), Headers: r.Header, Body: string(body)}
		ds.InsertRecord(record)

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	}
}

func GetRecordsHandler(ds *InMemoryDataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		records := ds.GetRecords()

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")

		t, err := template.ParseFiles("./request_layout.html")
		if err != nil {
			w.WriteHeader(500)
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
		}

		t.Execute(w, records)
	}
}
