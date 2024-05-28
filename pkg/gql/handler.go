package gql

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	adminKey = "admin" // может писать комментарии к постам
	userKey  = "user"  // не может писать комментарии
)

func Handler(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Access-Key")
	if key != adminKey && key != userKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	var params struct {
		Query string `json:"query"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if key == adminKey {
		result := executeQuery(params.Query)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	} else {
		if isQuery(params.Query) {
			result := executeQuery(params.Query)

			if err := json.NewEncoder(w).Encode(result); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}
	}
}

func isQuery(query string) bool {
	return strings.Contains(strings.TrimSpace(query), "posts")
}
