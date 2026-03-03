package services

import (
	"sync"
	"testing"
	"sistema-gestion-libros-electronicos-go2/models"
)

func TestConcurrentBookCreation(t *testing.T) {
	service := NewBookService()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			book := models.Book{
				ID:     id,
				Title:  "Libro",
				Author: "Autor",
			}
			service.CreateBook(book)
		}(i)
	}

	wg.Wait()

	books := service.GetAllBooks()
	if len(books) != 50 {
		t.Errorf("Se esperaban 50 libros, se obtuvieron %d", len(books))
	}
}
