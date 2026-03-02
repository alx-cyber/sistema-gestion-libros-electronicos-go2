package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/interfaces"
)

type LibroHandler struct {
	Gestor interfaces.GestorLibros
}

func (h *LibroHandler) BuscarLibro(w http.ResponseWriter, r *http.Request) {

	id := 1 // por ahora fijo para probar

	libro, err := h.Gestor.BuscarLibro(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libro)
}
