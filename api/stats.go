package api

import (
	"encoding/json"
	"net/http"
	"spellcheck-service/spellcheck_service"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	stats := spellcheck_service.GetStatistics()
	json.NewEncoder(w).Encode(stats)
}
