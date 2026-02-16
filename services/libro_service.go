package services

import (
	"errors"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"
)

// SistemaGestion administra los libros usando un map.
type SistemaGestion struct {
	libros map[int]*models.Libro
}

// NuevoSistema crea una nueva instancia del sistema.
func NuevoSistema() *SistemaGestion {
	return &SistemaGestion{
		libros: make(map[int]*models.Libro),
	}
}

// AgregarLibro agrega un libro al sistema.
func (s *SistemaGestion) AgregarLibro(libro *models.Libro) error {

	if _, existe := s.libros[libro.ID()]; existe {
		return errors.New("el libro ya existe")
	}

	s.libros[libro.ID()] = libro
	return nil
}

// BuscarLibro busca un libro por su ID.
func (s *SistemaGestion) BuscarLibro(id int) (*models.Libro, error) {

	libro, existe := s.libros[id]

	if !existe {
		return nil, errors.New("libro no encontrado")
	}

	return libro, nil
}
