package main

import (
	"log"
	"net/http"

	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/handlers"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/interfaces"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/services"
)

func main() {

	var gestor interfaces.GestorLibros = services.NuevoSistema()

	handler := &handlers.LibroHandler{
		Gestor: gestor,
	}

	http.HandleFunc("/libros", handler.ListarLibros)
	http.HandleFunc("/libro", handler.BuscarLibro)
	http.HandleFunc("/crear", handler.CrearLibro)

	log.Println("Servidor ejecutándose en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
