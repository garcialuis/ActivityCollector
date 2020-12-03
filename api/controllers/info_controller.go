package controllers

import (
	"encoding/json"
	"net/http"
)

func (server *Server) Info(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ActivityCollector /info")
}
