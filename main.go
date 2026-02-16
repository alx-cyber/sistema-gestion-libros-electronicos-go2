package main

import (
	"fmt"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/interfaces"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/models"
	"github.com/alx-cyber/sistema-gestion-libros-electronicos-go2/services"
)

func main() {

	// Usamos la interfaz como tipo
	var gestor interfaces.GestorLibros = services.NuevoSistema()

	// Crear libro usando constructor
	libro1, err := models.NuevoLibro(1, "Programacion en Go", "Alex")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Agregar libro
	err = gestor.AgregarLibro(libro1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Buscar libro
	libroEncontrado, err := gestor.BuscarLibro(1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Libro encontrado:", libroEncontrado.Titulo())
}
