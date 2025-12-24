package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ggarnhart/gogogo/internal/database"
	"github.com/ggarnhart/gogogo/internal/models"
)

type RequestHandler struct {
	db *database.DB
}

func NewRequestHandler(db *database.DB) *RequestHandler {
	return &RequestHandler{db: db}
}

func (h *RequestHandler) CreateRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var requestData models.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	request := models.CreateRequest{
		Message:   requestData.Message,
		IPAddress: requestData.IPAddress,
	}

	h.db.CreateRequest(r.Context(), &request)

	w.WriteHeader(http.StatusCreated)
}

func (h *RequestHandler) GetRequestsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	page := 0
	if v := r.URL.Query().Get("page"); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			if p > 0 {
				page = p
			}
		}
	}

	response, err := h.db.GetRequests(r.Context(), page)

	if err != nil {
		// return error
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
