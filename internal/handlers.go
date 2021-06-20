package internal

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

func CreateRecordHandler(ds *DataStore) func(http.ResponseWriter, *http.Request) {
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

func GetRecordsHandler(ds *DataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// We want to get records in reverse order.
		// This solution come from the official Go wiki.
		// https://github.com/golang/go/wiki/SliceTricks#reversing
		records := ds.GetRecords()
		for i := len(records)/2 - 1; i >= 0; i-- {
			opp := len(records) - 1 - i
			records[i], records[opp] = records[opp], records[i]
		}

		templateFile, err := getTemplatePath()
		if err != nil {
			w.WriteHeader(500)
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
		}

		t, err := template.ParseFiles(templateFile)
		if err != nil {
			w.WriteHeader(500)
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")

		t.Execute(w, records)
	}
}

func getTemplatePath() (string, error) {
	if _, err := os.Stat("./request_layout.html"); err == nil {
		return "./request_layout.html", nil
	}

	if _, err := os.Stat("../request_layout.html"); err == nil {
		return "../request_layout.html", nil
	}

	return "", fmt.Errorf("unable to find an HTML layout file")
}
