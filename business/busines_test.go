package business

import (
	"fmt"
	"log"
	"testing"
	"travelling-routes/dtstructs"
	"travelling-routes/utils"
)

func TestFindMinorCostRouteBetweenGRUandCDGFromCSVFile(t *testing.T) {

	wantedPrice, wantedRoute := float32(40), []string{"GRU", "BRC", "SCL", "ORL", "CDG"}

	b := Business{}

	gottemPrice, gottemRoute, err := b.RetrieveMinorCostRouteFromCSV("../input-routes.csv", "GRU", "CDG")
	if err != nil {
		t.Errorf("got error %v; wantted %.2f as price and %v as route", err, wantedPrice, wantedRoute)
	}

	if (wantedPrice != gottemPrice) && (utils.Equal(wantedRoute, gottemRoute)) {
		t.Errorf("got %.2f as price and %v as route; wantted %.2f as price and %v as route", gottemPrice, gottemRoute, wantedPrice, wantedRoute)
	}

	fmt.Printf("\n%.2f %v\n", gottemPrice, gottemRoute)
}

func TestErrorToFindMinorCostRouteBetweenXXXandXXXFromCSVFile(t *testing.T) {

	b := Business{}
	_, _, err := b.RetrieveMinorCostRouteFromCSV("../input-routes.csv", "XXX", "CDG")
	log.Println(err)
	if err == nil {
		t.Errorf("it got no error; should got 'origin [XXX] or destination [CDG] is not valid' error")
	}

	_, _, err = b.RetrieveMinorCostRouteFromCSV("../input-routes.csv", "GRU", "XXX")
	log.Println(err)
	if err == nil {
		t.Errorf("it got no error; should got 'origin [GRU] or destination [XXX] is not valid' error")
	}
}

func TestUpdateCSVfromGraph(t *testing.T) {

	filepath := "../../input-routes.csv"
	filepath_new := "../input-routes-new.csv"
	b := Business{}
	graph, err := b.BuildGraphFromCSV(filepath)
	if err != nil {
		t.Errorf("got err %v, expected to receive a graph", err)
	}

	got := graph.Exists("ZZZ")
	if got {
		t.Errorf("got: %v, wanted to be false [airport does not exists]", got)
	}

	graph.AddEdge("ZZZ", "CHI", float32(333.33))
	graph.AddEdge("GRU", "BRC", float32(10)) // <- should not be added again

	log.Printf("Graph before: %v", graph)

	err = b.UpdateCSVfromGraph(filepath_new, graph)
	if err != nil {
		t.Errorf("got err %v", err)
	}

	graph, err = b.BuildGraphFromCSV(filepath_new)
	if err != nil {
		t.Errorf("got err %v, expected to receive a graph", err)
	}

	got = graph.Exists("ZZZ")
	if !got {
		t.Errorf("got: %v, wanted to be true [airport ZZZ should exists]", got)
	}

	log.Printf("Graph after: %v", graph)

	err = b.UpdateCSVfromGraph(filepath_new, dtstructs.NewGraph())
	if err != nil {
		t.Errorf("got err %v", err)
	}

	graph, err = b.BuildGraphFromCSV(filepath_new)
	if err != nil {
		t.Errorf("got err %v, expected to receive a graph", err)
	}

	log.Printf("Graph final: %v", graph)
}
