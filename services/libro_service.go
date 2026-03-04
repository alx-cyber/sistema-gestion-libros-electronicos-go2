package services

import (
	"errors"
	"sync"

	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"
)

type SistemaGestion struct {
	libros map[int]*models.Libro
	mu     sync.Mutex
}

func NuevoSistema() *SistemaGestion {
	return &SistemaGestion{
		libros: make(map[int]*models.Libro),
	}
}

func (s *SistemaGestion) AgregarLibro(libro *models.Libro) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.libros[libro.ID]; existe {
		return errors.New("el libro ya existe")
	}

	s.libros[libro.ID] = libro
	return nil
}

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

func (s *SistemaGestion) ListarLibros() []*models.Libro {
	s.mu.Lock()
	defer s.mu.Unlock()

	var lista []*models.Libro
	for _, libro := range s.libros {
		lista = append(lista, libro)
	}
	return lista
}

func (s *SistemaGestion) ActualizarLibro(libro *models.Libro) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.libros[libro.ID]; !existe {
		return errors.New("libro no existe")
	}

	s.libros[libro.ID] = libro
	return nil
}

func (s *SistemaGestion) EliminarLibro(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.libros[id]; !existe {
		return errors.New("libro no encontrado")
	}

	delete(s.libros, id)
	return nil
}

func (s *SistemaGestion) TotalLibros() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.libros)
}
