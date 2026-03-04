package models

import "errors"

type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Anio   int    `json:"anio"`
	ISBN   string `json:"isbn"`
}

func NuevoLibro(id int, titulo string, autor string, anio int, isbn string) (*Libro, error) {
	if titulo == "" || autor == "" {
		return nil, errors.New("titulo y autor no pueden estar vacíos")
	}

	return &Libro{
		ID:     id,
		Titulo: titulo,
		Autor:  autor,
		Anio:   anio,
		ISBN:   isbn,
	}, nil
}

