package localization

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
	r.Get("/localization/{contentId}", h.Get)
	r.Patch("/localization/{contentId}", h.Update)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	contentID := chi.URLParam(r, "contentId")
	lang := r.URL.Query().Get("lang")
	result, err := h.svc.GetByContentID(contentID, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	contentID := chi.URLParam(r, "contentId")
	var req UpdateLocalizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	l, err := h.svc.Upsert(contentID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, l)
}
