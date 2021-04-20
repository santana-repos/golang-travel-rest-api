package dtstructs

import (
	"log"
	"testing"
	"travelling-routes/csv"
)

func TestCSVRouteSet(t *testing.T) {
	want := 1

	set := NewCSVRouteSet()
	route := csv.CSVroute{Origin: "GRU", Destination: "BRC", Cost: float32(10)}
	set.Add(route)
	set.Add(route)
	set.Add(route)

	got2 := set.Contains(route)
	if !got2 {
		t.Errorf("Got: %t; wannted to be true [ route: %v should exists ]", got2, route)
	}

	got := set.NumElements()
	if got != want {
		t.Errorf("Set contains: %d; Should have: %d", got, want)
	}

	log.Printf("Set: %v", set)
}
