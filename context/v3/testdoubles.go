package context2

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}

		}
		data <- result

	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// SpyResponseWriter checks whether a response has been written.
type SpyResponseWriter struct {
	written bool
}

// Header will mark written to true.
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

// Write will mark written to true.
func (s *SpyResponseWriter) Write([]byte) (int, error) {

	s.written = true
	return 0, errors.New("not implemeneted")
}

// WriteHeader will mark written to true.
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
