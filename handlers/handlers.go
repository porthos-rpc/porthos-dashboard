package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/porthos-rpc/porthos-dashboard/storage"
)

// IndexHandler will display the dashboard index page.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

// NewMethodsHandler creates a new handler to return stats for metrics.
func NewMethodsHandler(storage storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sinceStr := r.URL.Query().Get("since")
		var since time.Duration

		if sinceStr != "" {
			since, _ = time.ParseDuration(sinceStr)
		} else {
			since = -30 * time.Minute
		}

		metrics, err := storage.FindMethodMetrics(time.Now().Add(since))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(metrics)

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
