package dtstructs

import (
	"fmt"
	"testing"
)

func addGRUStretch(heap *heap) (*heap, Route) {
	route := Route{Price: 10, Nodes: []string{"GRU", "BRC"}}
	heap.Push(route)

	return heap, route
}

func TestPushARouteIntoTheHeap(t *testing.T) {
	heap := NewHeap()

	got := heap.Len()

	if got != 0 {
		t.Errorf("got %d, wanted 0", got)
	}

	heap, _ = addGRUStretch(heap)

	want := 1
	got = heap.Len()

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPushAndPopGRUBRCRouteFromTheHeap(t *testing.T) {
	heap := NewHeap()

	got := heap.Len()

	if got != 0 {
		t.Errorf("got %d, wanted 0", got)
	}

	addGRUStretch(heap)
	route2 := Route{Price: 20, Nodes: []string{"SCL", "ORL"}}
	heap.Push(route2)

	want := 2
	got = heap.Len()

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	// expected to take the route with minor price
	gottemRoute := heap.Pop()
	if (gottemRoute.Price != 10) && (gottemRoute.Nodes[0] != "GRU") && (gottemRoute.Nodes[1] != "BRC") {
		t.Errorf("got %v, wanted {10 [GRU BRC]}", got)
	}
	fmt.Printf("Removed Route: %v", gottemRoute)

	want = 1
	got = heap.Len()

	if got != want {
		t.Errorf("got %d, wanted %d after to pop a item from the Heap", got, want)
	}
}
