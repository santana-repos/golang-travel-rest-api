package dtstructs

import (
	"sort"
	"travelling-routes/csv"
)

// Implementação de Grafo para organizar os trechos
// das rotas dos aeroportos e seus correspondentes custos.
// Através desse Grafo os trechos de viagens são representados
// por arestas (edges) e inseridos na Árvore Heap, que por sua vez
// faz a ordenação dos trechos baseado no preço do trecho

type edge struct {
	node string
	cost float32
}

type Graph struct {
	nodes    map[string][]edge
	routeset routeSet
}

// NewGraph constructs a Graph to holds the representation of the airports
// as nodes and the cost/distance between them as edges.
// It offers the AddEdge function to help us to add two nodes and
// an Edge between them.
// Moreover, the graph offers the GetMinorPriceRoute function that
// receive two airports as origin and destiny parameters to return
// the minor price route to achieve the destination from the origin airport.
// It also initialize the Route Set
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]edge), routeset: NewCSVRouteSet()}
}

// AddEdge It offers the AddEdge function to help us to add two Nodes (origin
// and destiny as Strings) and an Edge (cost as Float32) between them.
func (g *Graph) AddEdge(origin, destiny string, cost float32) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, cost: cost})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, cost: cost})

	g.routeset.Add(csv.RouteData{Origin: origin, Destination: destiny, Cost: cost})
}

func (g *Graph) getEdge(node string) []edge {
	return g.nodes[node]
}

// GetMinorCostRoute function that receive two airports as origin and destiny
// parameters to return the minor price route to achieve the destination (destiny
// as string) from the origin (origin as string) airport.
// This functions implements the Dijkstra's algorithm as strategy to find
// the minor price route from origin to destiny.
func (g *Graph) GetMinorCostRoute(origin, destiny string) (float32, []string) {
	heap := NewHeap()
	heap.Push(Route{Price: float32(0), Nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*heap.prices) > 0 {
		// Find the costless yet to visit node
		p := heap.Pop()
		node := p.Nodes[len(p.Nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.Price, p.Nodes
		}

		for _, e := range g.getEdge(node) {
			if !visited[e.node] {
				// it calculates the total spent so far plus the cost and the route of getting here
				heap.Push(Route{Price: p.Price + e.cost, Nodes: append([]string{}, append(p.Nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

func (g *Graph) GetAllAirportsCodes(sortItens bool) []string {
	keys := make([]string, 0, len(g.nodes))
	for k := range g.nodes {
		keys = append(keys, k)
	}

	if sortItens {
		sort.Strings(keys)
	}

	return keys
}

func (g *Graph) Exists(airportName string) bool {
	exists := false

	if len(airportName) < 3 {
		return exists
	}

	airportNames := g.GetAllAirportsCodes(false)

	for _, name := range airportNames {
		if name == airportName {
			exists = true
		}
	}

	return exists
}

func (g *Graph) GetGraphAllRoutes() *[]csv.RouteData {
	routes := g.routeset.GetItems()

	return &routes
}
