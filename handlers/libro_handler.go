package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/interfaces"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"
)

type LibroHandler struct {
	Gestor interfaces.GestorLibros
}

// GET /libros
func (h *LibroHandler) ListarLibros(w http.ResponseWriter, r *http.Request) {

	libros := h.Gestor.ListarLibros()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}

// GET /libro?id=1
func (h *LibroHandler) BuscarLibro(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	libro, err := h.Gestor.BuscarLibro(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libro)
}

// POST /libros
func (h *LibroHandler) CrearLibro(w http.ResponseWriter, r *http.Request) {

	var datos struct {
		ID     int    `json:"id"`
		Titulo string `json:"titulo"`
		Autor  string `json:"autor"`
	}

	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	libro, err := models.NuevoLibro(datos.ID, datos.Titulo, datos.Autor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Gestor.AgregarLibro(libro)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *LibroHandler) ActualizarLibro(w http.ResponseWriter, r *http.Request) {

	var datos struct {
		ID     int    `json:"id"`
		Titulo string `json:"titulo"`
		Autor  string `json:"autor"`
	}

	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	libro, err := models.NuevoLibro(datos.ID, datos.Titulo, datos.Autor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Gestor.ActualizarLibro(libro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *LibroHandler) EliminarLibro(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.Gestor.EliminarLibro(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	func (h *LibroHandler) TotalLibros(w http.ResponseWriter, r *http.Request) {

	total := h.Gestor.TotalLibros()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"total": total})
}

	w.WriteHeader(http.StatusOK)
}
func (h *LibroHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Servidor funcionando correctamente"))
}
func (h *LibroHandler) PruebaConcurrente(w http.ResponseWriter, r *http.Request) {

	ch := make(chan string)

	libro1, _ := models.NuevoLibro(100, "Libro Concurrente 1", "Autor X")
	libro2, _ := models.NuevoLibro(101, "Libro Concurrente 2", "Autor Y")

	go h.Gestor.AgregarLibroConcurrente(libro1, ch)
	go h.Gestor.AgregarLibroConcurrente(libro2, ch)

	msg1 := <-ch
	msg2 := <-ch

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{msg1, msg2})
}
