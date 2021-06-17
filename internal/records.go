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
	M       sync.Mutex
	records []Record
}

func (ds *InMemoryDataStore) InsertRecord(r Record) {
	ds.M.Lock()
	defer ds.M.Unlock()

	ds.records = append(ds.records, r)
}
func (ds *InMemoryDataStore) GetRecords() []Record {
	return ds.records
}
