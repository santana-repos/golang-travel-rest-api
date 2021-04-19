package dtstructs

import (
	"fmt"
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func buildBaseGraph() *graph {
	graph := NewGraph()
	graph.AddEdge("GRU", "BRC", 10)
	graph.AddEdge("BRC", "SCL", 5)
	graph.AddEdge("GRU", "CDG", 75)
	graph.AddEdge("GRU", "SCL", 20)
	graph.AddEdge("GRU", "ORL", 56)
	graph.AddEdge("ORL", "CDG", 5)
	graph.AddEdge("SCL", "ORL", 20)

	return graph
}

func TestFindMinorPriceRouteBetweenGRUandCDG(t *testing.T) {
	wantedPrice, wantedRoute := float32(40), []string{"GRU", "BRC", "SCL", "ORL", "CDG"}

	graph := buildBaseGraph()
	gottemPrice, gottemRoute := graph.GetMinorPriceRoute("GRU", "CDG")

	if (wantedPrice != gottemPrice) && (equal(wantedRoute, gottemRoute)) {
		t.Errorf("got %.2f as price and %v as route; wantted %.2f as price and %v as route", gottemPrice, gottemRoute, wantedPrice, wantedRoute)
	}

	fmt.Printf("\n%.2f %v\n", gottemPrice, gottemRoute)
}

func TestFindMinorPriceRouteBetweenBRCandORL(t *testing.T) {
	wantedPrice, wantedRoute := float32(25), []string{"BRC", "SCL", "ORL"}

	graph := buildBaseGraph()
	gottemPrice, gottemRoute := graph.GetMinorPriceRoute("BRC", "ORL")

	if (wantedPrice != gottemPrice) && (equal(wantedRoute, gottemRoute)) {
		t.Errorf("got %.2f as price and %v as route; wantted %.2f as price and %v as route", gottemPrice, gottemRoute, wantedPrice, wantedRoute)
	}

	fmt.Printf("\n%.2f %v\n", gottemPrice, gottemRoute)
}
