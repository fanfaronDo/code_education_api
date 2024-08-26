package handler

import (
	"encoding/json"
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"net/http"
)

func (h *Handler) getNotes(w http.ResponseWriter, r *http.Request) {
	notesData := struct {
		Data []domain.Note `json:"data"`
	}{}

	userID, err := getUserId(r)
	if err != nil {
		http.Error(w, "User not find in context", http.StatusBadRequest)
		return
	}
	notesData.Data = h.service.NotesService.GetNotes(userID)
	if err := json.NewEncoder(w).Encode(notesData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
