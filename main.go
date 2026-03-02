package main

import (
	"log"
	"net/http"

	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/handlers"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/interfaces"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/services"
)

func main() {

	var gestor interfaces.GestorLibros = services.NuevoSistema()

	// Crear libro de prueba
	libro1, _ := models.NuevoLibro(1, "Programacion en Go", "Alex")
	gestor.AgregarLibro(libro1)

	handler := &handlers.LibroHandler{
		Gestor: gestor,
	}

	http.HandleFunc("/libro", handler.BuscarLibro)

	log.Println("Servidor ejecutándose en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
