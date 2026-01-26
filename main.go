

package main

import (
	"fmt"
	"sistema/models"
	"sistema/services"
)

func main() {
	var opcion int
	libros := []models.Libro{}

	for {
		fmt.Println("\n--- Sistema de Gestión de Libros Electrónicos ---")
		fmt.Println("1. Registrar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libro")
		fmt.Println("4. Eliminar libro")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			libros = services.RegistrarLibro(libros)
		case 2:
			services.ListarLibros(libros)
		case 3:
			services.BuscarLibro(libros)
		case 4:
			libros = services.EliminarLibro(libros)
		case 5:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
