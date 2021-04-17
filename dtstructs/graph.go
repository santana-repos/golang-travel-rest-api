package dtstructs

type stretch struct {
	node string
	cost float32
}

type graph struct {
	nodes map[string][]stretch
}

func NewGraph() *graph {
	return &graph{nodes: make(map[string][]stretch)}
}

func (g *graph) AddStretch(origin, destiny string, custo float32) {
	g.nodes[origin] = append(g.nodes[origin], stretch{node: destiny, cost: custo})
	g.nodes[destiny] = append(g.nodes[destiny], stretch{node: origin, cost: custo})
}

func (g *graph) getStretch(node string) []stretch {
	return g.nodes[node]
}

func (g *graph) GetMinorPriceRoute(origin, destiny string) (float32, []string) {
	heap := NewHeap()
	heap.Push(Route{Price: float32(0), Nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*heap.prices) > 0 {
		// Find the nearest yet to visit node
		p := heap.Pop()
		node := p.Nodes[len(p.Nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.Price, p.Nodes
		}

		for _, e := range g.getStretch(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				heap.Push(Route{Price: p.Price + e.cost, Nodes: append([]string{}, append(p.Nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
