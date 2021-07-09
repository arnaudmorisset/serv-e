package internal

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"serv-e/pkg"
	"time"
)

var POSSIBLE_TEMPLATE_LOCATIONS = []string{"./request_layout.html", "../request_layout.html"}

func CreateRecordHandler(ds *DataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			// TODO: improve with a proper log handler ?
			fmt.Fprintf(os.Stderr, "cannot read body: %s", err.Error())
			returnErrorResponse(w, pkg.ErrCannotReadBody)
			return
		}

		record := Record{Id: time.Now().Format("15:04:05"), Headers: r.Header, Body: string(body)}
		ds.InsertRecord(record)

		returnCreatedRecordResponse(w)
	}
}

func GetRecordsHandler(ds *DataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// We want to get records in reverse order.
		records := ds.GetRecords()
		ReverseRecords(&records)

		templateFile, err := getTemplatePath()
		if err != nil {
			fmt.Fprintf(os.Stderr, "no template found")
			returnErrorResponse(w, err)
			return
		}

		t, err := template.ParseFiles(templateFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot parse template: %s", err.Error())
			returnErrorResponse(w, pkg.ErrCannotParseTemplate)
			return
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")

		if err := t.Execute(w, records); err != nil {
			fmt.Fprintf(os.Stderr, "cannot execute template: %s", err.Error())
			returnErrorResponse(w, pkg.ErrTemplateExecution)
			return
		}
	}
}

func getTemplatePath() (string, error) {
	for _, possibleLocation := range POSSIBLE_TEMPLATE_LOCATIONS {
		if _, err := os.Stat(possibleLocation); err == nil {
			return possibleLocation, nil
		}
	}

	return "", pkg.ErrCannotFindHTMLLayoutFile
}
