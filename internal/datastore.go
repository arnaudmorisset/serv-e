package internal

import (
	"net/http"
	"sync"
)

type Record struct {
	Id      string
	Headers http.Header
	Body    string
}

type DataStore struct {
	m       sync.Mutex
	records []Record
}

func (ds *DataStore) InsertRecord(r Record) {
	ds.m.Lock()
	defer ds.m.Unlock()

	ds.records = append(ds.records, r)
}
func (ds *DataStore) GetRecords() []Record {
	return ds.records
}
