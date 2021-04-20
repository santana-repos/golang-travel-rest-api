package business

import (
	"fmt"
	"strconv"
	"travelling-routes/csv"
	"travelling-routes/dtstructs"
)

func BuildGraphFromCSV(filepath string) (*dtstructs.Graph, error) {

	graph := dtstructs.NewGraph()

	lines, err := csv.LoadCSVlines(filepath)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		price, err := strconv.ParseFloat(line[2], 32)
		if err != nil {
			return nil, err
		}

		graph.AddEdge(line[0], line[1], float32(price))
	}

	return graph, err
}

func UpdateCSVfromGraph(filepath string, graph *dtstructs.Graph) error {

	//csvroutes := make([]csv.CSVroute, 0, 10)

	/*
		for _,

		graph := dtstructs.NewGraph()

		lines, err := csv.LoadCSVlines(filepath)
		if err != nil {
			return nil, err
		}

		for _, line := range lines {
			price, err := strconv.ParseFloat(line[2], 32)
			if err != nil {
				return nil, err
			}

			graph.AddEdge(line[0], line[1], float32(price))
		}
	*/

	return nil
}

func RetrieveMinorCostRouteFromCSV(filepath string, origin string, destination string) (float32, []string, error) {
	graph, err := BuildGraphFromCSV(filepath)
	if err != nil {
		return 0, nil, err
	}

	if (!graph.Exists(origin)) || (!graph.Exists(destination)) {
		return 0, nil, fmt.Errorf("origin [%s] or destination [%s] is not valid", origin, destination)
	}

	cost, route := graph.GetMinorCostRoute(origin, destination)

	return cost, route, nil
}
