package context

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			slog.Error("fetching data", "error", err)
			return
		}
		fmt.Fprint(w, data)
	}
}
