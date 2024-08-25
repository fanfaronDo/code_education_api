package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h3>SIGN UP</h3>")
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("This is sign in"))
}
