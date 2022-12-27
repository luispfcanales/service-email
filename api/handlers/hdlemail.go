package handlers

import (
	"encoding/json"
	"net/http"
)

// SendEmail send email
func SendEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	p := &Person{}
	json.NewDecoder(r.Body).Decode(p)
	notifyByEmail("notify.html", p)
	json.NewEncoder(w).Encode(p)
}
