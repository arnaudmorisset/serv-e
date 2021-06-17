package main

import (
	"fmt"
	"net/http"
	"os"
	"serv-e/internal"
)

func main() {
	var ds internal.InMemoryDataStore

	createRecordHandler := internal.CreateRecordHandler{Ds: &ds}
	getRecordsHandler := internal.GetRecordsHandler{Ds: &ds}

	http.HandleFunc("/", createRecordHandler.ServeHTTP)
	http.HandleFunc("/records", getRecordsHandler.ServeHTTP)

	fmt.Println("server listening on http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Fprintf(os.Stderr, "server closed: %v", err)
		os.Exit(1)
	}
}
