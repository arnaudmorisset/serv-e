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

type InMemoryDataStore struct {
	m       sync.Mutex
	records []Record
}

func (ds *InMemoryDataStore) InsertRecord(r Record) {
	ds.m.Lock()
	defer ds.m.Unlock()

	ds.records = append(ds.records, r)
}
func (ds *InMemoryDataStore) GetRecords() []Record {
	return ds.records
}
