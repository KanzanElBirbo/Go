package main

import (
"fmt"
"math"
)

// INF představuje hodnotu nekonečna pro inicializaci vzdáleností
var INF = math.Inf(1)

// dijkstra je funkce implementující Dijkstrův algoritmus
func dijkstra(start string, graph map[string]map[string]int) map[string]map[string]interface{} {
	// Inicializace map pro vzdálenosti, předchozí uzly a navštívené uzly
	distances := make(map[string]float64)
	previous := make(map[string]string)
	visited := make(map[string]bool)

	// Inicializace vzdáleností na nekonečno a předchozích uzlů na prázdný řetězec
	for node := range graph {
		distances[node] = INF
		previous[node] = ""
	}

	// Nastavení vzdálenosti od startovního uzlu na 0
	distances[start] = 0

	// Hlavní část algoritmu
	for len(visited) < len(graph) {
		// Najdi uzel s nejmenší vzdáleností
		minNode := minDistance(distances, visited)

		// Procházej všechny sousedy tohoto uzlu
		for neighbor, distance := range graph[minNode] {
			// Spočítej alternativní vzdálenost
			alt := distances[minNode] + float64(distance)

			// Pokud je alternativní vzdálenost menší než aktuální vzdálenost, aktualizuj hodnoty
			if alt < distances[neighbor] {
				distances[neighbor] = alt
				previous[neighbor] = minNode
			}
		}

		// Přidej tento uzel do navštívených
		visited[minNode] = true
	}

	// Vytvoř výsledný seznam s informacemi o vzdálenostech a předchozích uzlech
	result := make(map[string]map[string]interface{})
	for node, distance := range distances {
		result[node] = map[string]interface{}{
			"vzdalenost": distance,
			"predchozi":  previous[node],
		}
	}

	return result
}

// minDistance najde uzel s nejmenší vzdáleností
func minDistance(distances map[string]float64, visited map[string]bool) string {
	minDist := INF
	var minNode string

	// Prochází všechny uzly a hledá ten s nejmenší vzdáleností, který ještě nebyl navštíven
	for node, dist := range distances {
		if dist <= minDist && !visited[node] {
			minDist = dist
			minNode = node
		}
	}

	return minNode
}

func main() {
	// Příklad grafu
	graph1 := map[string]map[string]int{
		"A": {"B": 2, "C": 4},
		"B": {"A": 2, "C": 1, "F": 5},
		"C": {"A": 4, "D": 3},
		"D": {"C": 3, "A": 1},
		"E": {},
		"F": {"B": 5},
	}

	// Počáteční uzel pro Dijkstrův algoritmus
	startNode := "A"

	// Zavolání Dijkstrova algoritmu a získání výsledků
	result := dijkstra(startNode, graph1)

	// Výpis výsledků
	fmt.Printf("Nejkratší vzdálenosti od uzlu %s:\n", startNode)
	for node, info := range result {
		fmt.Printf("%s: vzdálenost = %v, předchozí = %v\n", node, info["vzdalenost"], info["predchozi"])
	}
}
