package interfaces

import "github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"

// GestorLibros define el comportamiento del sistema de gestión.
type GestorLibros interface {
	AgregarLibro(libro *models.Libro) error
	BuscarLibro(id int) (*models.Libro, error)
	ListarLibros() []*models.Libro
	ActualizarLibro(libro *models.Libro) error
	EliminarLibro(id int) error
	AgregarLibroConcurrente(libro *models.Libro, ch chan string)
	TotalLibros() int
}
