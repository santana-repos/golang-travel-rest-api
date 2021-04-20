package csv

import (
	"testing"
)

func TestLoadCSVroutes(t *testing.T) {
	filepath := "../input-routes-test.csv"
	routes := make([]RouteData, 0, 10)
	routes = append(routes, RouteData{Origin: "GRU", Destination: "BRC", Cost: float32(10)})
	routes = append(routes, RouteData{Origin: "GRU", Destination: "BRC", Cost: float32(10)})
	routes = append(routes, RouteData{Origin: "BRC", Destination: "SCL", Cost: float32(5)})
	routes = append(routes, RouteData{Origin: "GRU", Destination: "CDG", Cost: float32(75)})
	routes = append(routes, RouteData{Origin: "GRU", Destination: "SCL", Cost: float32(20)})
	routes = append(routes, RouteData{Origin: "GRU", Destination: "ORL", Cost: float32(56)})
	routes = append(routes, RouteData{Origin: "ORL", Destination: "CDG", Cost: float32(5)})
	routes = append(routes, RouteData{Origin: "SCL", Destination: "ORL", Cost: float32(20)})

	err := WriteCsv(filepath, routes, true)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
}
