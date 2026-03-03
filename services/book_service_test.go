package services

import (
	"testing"
	"sistema-gestion-libros-electronicos-go2/models"
)

func TestCreateBook(t *testing.T) {
	service := NewBookService()

	book := models.Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Autor Test",
	}

	err := service.CreateBook(book)
	if err != nil {
		t.Errorf("Error inesperado al crear libro: %v", err)
	}

	books := service.GetAllBooks()
	if len(books) != 1 {
		t.Errorf("Se esperaba 1 libro, se obtuvo %d", len(books))
	}
}

func TestDeleteBook(t *testing.T) {
	service := NewBookService()

	book := models.Book{
		ID:     1,
		Title:  "Libro eliminar",
		Author: "Autor",
	}

	service.CreateBook(book)
	err := service.DeleteBook(1)

	if err != nil {
		t.Errorf("No se pudo eliminar libro: %v", err)
	}

	books := service.GetAllBooks()
	if len(books) != 0 {
		t.Errorf("El libro no fue eliminado correctamente")
	}
}
