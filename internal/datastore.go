package internal

import (
	"net/http"
	"reflect"
	"sync"
)

type Record struct {
	Id      string
	Headers http.Header
	Body    string
}

func (r Record) Equals(r2 Record) bool {
	return r.Id == r2.Id && reflect.DeepEqual(r.Headers, r2.Headers) && r.Body == r2.Body
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
