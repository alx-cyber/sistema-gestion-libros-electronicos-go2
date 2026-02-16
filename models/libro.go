package models

import "errors"

// Libro representa un libro electrónico.
// Usamos campos privados para encapsulación.
type Libro struct {
	titulo string
	autor  string
	anio   int
	isbn   string
}

// Constructor con validaciones.
func NuevoLibro(titulo string, autor string, anio int, isbn string) (*Libro, error) {
	if titulo == "" || autor == "" {
		return nil, errors.New("titulo y autor no pueden estar vacíos")
	}
	return &Libro{titulo: titulo, autor: autor, anio: anio, isbn: isbn}, nil
}

// Getters públicos:
func (l *Libro) Titulo() string { return l.titulo }
func (l *Libro) Autor() string  { return l.autor }
func (l *Libro) Anio() int      { return l.anio }
func (l *Libro) ISBN() string   { return l.isbn }

