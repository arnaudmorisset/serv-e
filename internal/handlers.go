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

		fmt.Printf("%+v", ds)

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
			returnErrorResponse(w, err)
			return
		}

		t, err := template.ParseFiles(templateFile)
		if err != nil {
			returnErrorResponse(w, err)
			return
		}

		returnOkResponse(w, "text/html", nil)

		t.Execute(w, records)
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
