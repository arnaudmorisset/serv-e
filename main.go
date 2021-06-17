package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"serv-e/internal"
	"time"
)

var ds internal.InMemoryDataStore

func main() {
	http.HandleFunc("/", createRecord)
	http.HandleFunc("/records", getRecords)

	fmt.Println("server listening on http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Fprintf(os.Stderr, "server closed: %v", err)
		os.Exit(1)
	}
}

func createRecord(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		sendServerError(writer, err)
	}

	record := internal.Record{Id: time.Now().Format("15:04:05"), Headers: request.Header, Body: string(body)}
	ds.InsertRecord(record)

	writer.WriteHeader(200)
	writer.Header().Add("Content-Type", "text/plain")
	writer.Write([]byte("OK"))
}

func getRecords(writer http.ResponseWriter, request *http.Request) {
	records := ds.GetRecords()

	writer.WriteHeader(200)
	writer.Header().Add("Content-Type", "text/html")

	t, err := template.ParseFiles("./request_layout.html")
	if err != nil {
		sendServerError(writer, err)
	}

	t.Execute(writer, records)
}

func sendServerError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(500)
	writer.Header().Add("Content-Type", "text/plain")
	writer.Write([]byte(err.Error()))
}
