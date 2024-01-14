package main

import "fmt" // Importuje základní balíček fmt

func main() { // Hlavní funkce programu
	// Zavolání funkce numbers
	numbers()

	// Zavolání funkce maximum
	maximum()
}

func numbers() { // Funkce pro inicializaci a výpis pole
	// Pole
	var arr [3]int                     // Deklarace pole celých čísel o velikosti 3
	arr[0] = 1                         // První prvek pole
	arr[1] = 2                         // Druhý prvek pole
	arr[2] = 3                         // Třetí prvek pole
	fmt.Println("Numbers array:", arr) // Výpis pole
}

func maximum() { // Funkce pro nalezení maximální hodnoty v poli
	numbers := [5]int{15, 6, 23, 8, 16} // Inicializace pole celých čísel o velikosti 5
	max := numbers[0]                   // Inicializace proměnné max prvním prvkem pole

	for _, value := range numbers { // Iterace přes prvky pole
		if value > max { // Porovnání aktuálního prvku s dosavadní maximální hodnotou
			max = value // Aktualizace maximální hodnoty, pokud je aktuální prvek větší
		}
	}

	fmt.Println("Maximum element:", max) // Výpis maximální hodnoty v poli
}
