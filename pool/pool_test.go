package pool

import (
	"sync"
	"testing"
	"time"
)

type Fields map[string]interface{}

type Entry struct {
	start time.Time

	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Fields    []Fields  `json:"fields"`
}

var responsePool = sync.Pool{
	New: func() interface{} {
		return &Entry{}
	},
}

func GetResponse() *Entry {
	return responsePool.Get().(*Entry)
}

func PutResponse(buf *Entry) {
	responsePool.Put(buf)
}

func BenchmarkCreateEntryWithPool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		response := GetResponse()
		response.start = time.Now()
		PutResponse(response)
	}
}

func BenchmarkCreateEntryWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		response := &Entry{}

		response.start = time.Now()

	}
}
