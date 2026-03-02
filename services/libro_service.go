package services

import (
	"errors"
	"sync"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"
)
// SistemaGestion administra los libros usando un map.
type SistemaGestion struct {
	libros map[int]*models.Libro
	mu     sync.Mutex
}

// NuevoSistema crea una nueva instancia del sistema.
func NuevoSistema() *SistemaGestion {
	return &SistemaGestion{
		libros: make(map[int]*models.Libro),
	}
}

// AgregarLibro agrega un libro al sistema.
func (s *SistemaGestion) AgregarLibro(libro *models.Libro) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.libros[libro.ID()]; existe {
		return errors.New("el libro ya existe")
	}

	s.libros[libro.ID()] = libro
	return nil
}

// BuscarLibro busca un libro por su ID.
func (s *SistemaGestion) BuscarLibro(id int) (*models.Libro, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	libro, existe := s.libros[id]

	if !existe {
		return nil, errors.New("libro no encontrado")
	}

	return libro, nil
}
func (s *SistemaGestion) AgregarLibroConcurrente(libro *models.Libro, ch chan string) {
	err := s.AgregarLibro(libro)

	if err != nil {
		ch <- "Error: " + err.Error()
		return
	}

	ch <- "Libro agregado correctamente"
}

