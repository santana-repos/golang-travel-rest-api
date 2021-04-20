package business

import (
	"fmt"
	"log"
	"testing"
	"travelling-routes/utils"
)

func TestFindMinorCostRouteBetweenGRUandCDGFromCSVFile(t *testing.T) {

	wantedPrice, wantedRoute := float32(40), []string{"GRU", "BRC", "SCL", "ORL", "CDG"}

	gottemPrice, gottemRoute, err := RetrieveMinorCostRouteFromCSV("../input-routes.csv", "GRU", "CDG")
	if err != nil {
		t.Errorf("got error %v; wantted %.2f as price and %v as route", err, wantedPrice, wantedRoute)
	}

	if (wantedPrice != gottemPrice) && (utils.Equal(wantedRoute, gottemRoute)) {
		t.Errorf("got %.2f as price and %v as route; wantted %.2f as price and %v as route", gottemPrice, gottemRoute, wantedPrice, wantedRoute)
	}

	fmt.Printf("\n%.2f %v\n", gottemPrice, gottemRoute)
}

func TestErrorToFindMinorCostRouteBetweenXXXandXXXFromCSVFile(t *testing.T) {

	_, _, err := RetrieveMinorCostRouteFromCSV("../input-routes.csv", "XXX", "CDG")
	log.Println(err)
	if err == nil {
		t.Errorf("it got no error; should got 'origin [XXX] or destination [CDG] is not valid' error")
	}

	_, _, err = RetrieveMinorCostRouteFromCSV("../input-routes.csv", "GRU", "XXX")
	log.Println(err)
	if err == nil {
		t.Errorf("it got no error; should got 'origin [GRU] or destination [XXX] is not valid' error")
	}
}

func TestUpdateCSVfromGraph(t *testing.T) {

	//graph, err := BuildGraphFromCSV("../input-routes-test.csv")
	graph, err := BuildGraphFromCSV("../input-routes-test.csv")
	if err != nil {
		t.Errorf("got err %v, expected to receive a graph", err)
	}

	/*for i, node := range graph {

	}*/
	log.Printf("Grafo: %v", graph)
}
