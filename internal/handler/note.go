package handler

import (
	"encoding/json"
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"net/http"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {

	userID, err := getUserId(r)
	var inputNote domain.Note
	if err := json.NewDecoder(r.Body).Decode(&inputNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.NoteService.CreateNote(userID, inputNote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"id": id,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
