package context2

import (
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch() string
	Cancel()
}

// Server returns a handler for calling Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string)
		go func() {
			data <- store.Fetch()
		}()
		// manually calling cancel if context is done
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
