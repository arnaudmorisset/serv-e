package main

import (
	"fmt"
	"net/http"
	"os"
	"serv-e/internal"
)

func main() {
	var ds internal.InMemoryDataStore

	http.HandleFunc("/", internal.CreateRecordHandler(&ds))
	http.HandleFunc("/records", internal.GetRecordsHandler(&ds))

	fmt.Println("server listening on http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Fprintf(os.Stderr, "server closed: %v", err)
		os.Exit(1)
	}
}
