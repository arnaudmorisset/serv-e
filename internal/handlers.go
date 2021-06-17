package internal

import (
	"html/template"
	"io"
	"net/http"
	"time"
)

type CreateRecordHandler struct {
	Ds *InMemoryDataStore
}

func (h *CreateRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendServerError(w, err)
	}

	record := Record{Id: time.Now().Format("15:04:05"), Headers: r.Header, Body: string(body)}
	h.Ds.InsertRecord(record)

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

type GetRecordsHandler struct {
	Ds *InMemoryDataStore
}

func (h *GetRecordsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	records := h.Ds.GetRecords()

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/html")

	t, err := template.ParseFiles("./request_layout.html")
	if err != nil {
		sendServerError(w, err)
	}

	t.Execute(w, records)
}

func sendServerError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(500)
	writer.Header().Add("Content-Type", "text/plain")
	writer.Write([]byte(err.Error()))
}
