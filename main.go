package main

import (
	"fmt"
	"streaming/catalog"
	"streaming/search"
)

func main() {
	fmt.Println("Catálogo de Películas Ecuatorianas")

	// Crear catálogo inicial
	catalogo := []catalog.Content{
		catalog.NewContent("Ratas, ratones y rateros", "Drama", 1999, 108),
		catalog.NewContent("Qué tan lejos", "Comedia dramática", 2006, 92),
		catalog.NewContent("Alba", "Drama", 2016, 94),
		catalog.NewContent("Sumergible", "Suspenso", 2020, 100),
		catalog.NewContent("Agujero negro", "Ciencia ficción", 2018, 95),
	}

	// Mostrar catálogo completo
	fmt.Println("\nCatálogo completo:")
	for _, c := range catalogo {
		c.Show()
	}

	// Buscar por género
	fmt.Println("\nContenido filtrado por género: Drama")
	resultado := search.FilterByGenre(catalogo, "Drama")
	for _, r := range resultado {
		r.Show()
	}

	// Agregar nuevo contenido validando datos
	fmt.Println("\nAgregando nuevo contenido con validación:")
	nuevo, err := catalog.NewContentSafe("El facilitador", "Drama político", 2014, 96)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		nuevo.Show()
	}
}
