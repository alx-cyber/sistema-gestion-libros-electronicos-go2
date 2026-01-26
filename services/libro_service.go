package services

import (
	"bufio"
	"fmt"
	"os"
	"sistema/models"
)

func RegistrarLibro(libros []models.Libro) []models.Libro {
	reader := bufio.NewReader(os.Stdin)

	var libro models.Libro

	fmt.Print("Título: ")
	libro.Titulo, _ = reader.ReadString('\n')

	fmt.Print("Autor: ")
	libro.Autor, _ = reader.ReadString('\n')

	fmt.Print("Año: ")
	fmt.Scanln(&libro.Anio)

	fmt.Print("ISBN: ")
	fmt.Scanln(&libro.ISBN)

	libros = append(libros, libro)
	fmt.Println("Libro registrado correctamente")

	return libros
}

func ListarLibros(libros []models.Libro) {
	if len(libros) == 0 {
		fmt.Println("No hay libros registrados")
		return
	}

	for i, libro := range libros {
		fmt.Printf("%d. %s - %s (%d)\n", i+1, libro.Titulo, libro.Autor, libro.Anio)
	}
}

func BuscarLibro(libros []models.Libro) {
	var titulo string
	fmt.Print("Ingrese el título a buscar: ")
	fmt.Scanln(&titulo)

	for _, libro := range libros {
		if libro.Titulo == titulo {
			fmt.Println("Libro encontrado:")
			fmt.Println("Título:", libro.Titulo)
			fmt.Println("Autor:", libro.Autor)
			fmt.Println("Año:", libro.Anio)
			fmt.Println("ISBN:", libro.ISBN)
			return
		}
	}
	fmt.Println("Libro no encontrado")
}

func EliminarLibro(libros []models.Libro) []models.Libro {
	var isbn string
	fmt.Print("Ingrese el ISBN del libro a eliminar: ")
	fmt.Scanln(&isbn)

	for i, libro := range libros {
		if libro.ISBN == isbn {
			libros = append(libros[:i], libros[i+1:]...)
			fmt.Println("Libro eliminado")
			return libros
		}
	}
	fmt.Println("Libro no encontrado")
	return libros
}


