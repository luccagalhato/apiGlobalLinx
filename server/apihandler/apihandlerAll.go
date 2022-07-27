package apiHandler

import (
	sql "MANCHESTER/API-GLOBAL-LINX/database"
	"encoding/json"
	"net/http"
)

// All ...
func All(s *sql.Str) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//data := sql.GetItem()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("access-control-expose-headers", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("data")
	}
}
