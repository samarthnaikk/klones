package availability

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/availability/{contentId}", h.Get)
	r.Patch("/availability/{contentId}", h.Update)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	contentID := chi.URLParam(r, "contentId")
	region := r.URL.Query().Get("region")
	result, err := h.svc.GetByContentID(contentID, region)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	contentID := chi.URLParam(r, "contentId")
	var req UpdateAvailabilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a, err := h.svc.Upsert(contentID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, a)
}
