package dtstructs

import (
	"fmt"
	"log"
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

func TestFindMinorCostRouteBetweenGRUandCDG(t *testing.T) {
	wantedPrice, wantedRoute := float32(40), []string{"GRU", "BRC", "SCL", "ORL", "CDG"}

	graph := buildBaseGraph()
	gottemPrice, gottemRoute := graph.GetMinorCostRoute("GRU", "CDG")

	if (wantedPrice != gottemPrice) && (equal(wantedRoute, gottemRoute)) {
		t.Errorf("got %.2f as price and %v as route; wantted %.2f as price and %v as route", gottemPrice, gottemRoute, wantedPrice, wantedRoute)
	}

	fmt.Printf("\n%.2f %v\n", gottemPrice, gottemRoute)
}

func TestFindMinorCostRouteBetweenBRCandORL(t *testing.T) {
	wantedPrice, wantedRoute := float32(25), []string{"BRC", "SCL", "ORL"}

	graph := buildBaseGraph()
	gottemPrice, gottemRoute := graph.GetMinorCostRoute("BRC", "ORL")

	if (wantedPrice != gottemPrice) && (equal(wantedRoute, gottemRoute)) {
		t.Errorf("got %.2f as price and %v as route; wantted %.2f as price and %v as route", gottemPrice, gottemRoute, wantedPrice, wantedRoute)
	}

	fmt.Printf("\n%.2f %v\n", gottemPrice, gottemRoute)
}

func TestFindAirportsNames(t *testing.T) {
	want := []string{"CDG", "ORL", "GRU", "BRC", "SCL"}

	graph := buildBaseGraph()
	got := graph.GetAllAirportsCodes(false)
	log.Printf("Nomes want: %v", want)
	log.Printf("Nomes got: %v", got)

	if !equal(got, want) {
		t.Errorf("got:\n%v\nwantted:\n%v", got, want)
	}
}

func TestFindAirportsNamesSorted(t *testing.T) {
	want := []string{"BRC", "CDG", "GRU", "ORL", "SCL"}

	graph := buildBaseGraph()
	got := graph.GetAllAirportsCodes(true)
	log.Printf("Nomes want: %v", want)
	log.Printf("Nomes got: %v", got)

	if !equal(got, want) {
		t.Errorf("got:\n%v\nwantted:\n%v", got, want)
	}
}

func TestFindAirportNameIsValid(t *testing.T) {
	graph := buildBaseGraph()
	got := graph.Exists("BRC")

	if !got {
		t.Errorf("got: %v, wanted to be true [airport exists]", got)
	}

	got = graph.Exists("XXX")

	if got {
		t.Errorf("got: %v, wanted to be false [airport does not exists]", got)
	}

	got = graph.Exists("BR")

	if got {
		t.Errorf("got: %v, wanted to be false [airport CODE should be invalid]", got)
	}
}
