package main

import (
	"fmt"
	"travelling-routes/dtstructs"
)

type Caminho struct {
	Origem    int
	Destino   int
	Distancia int
}
type Local struct {
	Codigo    int
	Sigla     string
	Descricao string
}

type Vertice struct {
	Local          Local
	MenorDistancia int
}

func main() {
	fmt.Println("hello Traveller o/")

	//caminhos := make([]Caminho, 0, 3)
	/*
		caminho1 := new(Caminho)
		caminho1.Trecho = "ITAP-BARUERI"
		caminho1.Distancia = 60
		caminhos = append(caminhos, *caminho1)
	*/
	/*
		local_GRU := Local{Codigo: 0, Sigla: "GRU", Descricao: "Aeroporto de Guarulhos-BRASIL"}
		local_BRC := Local{Codigo: 1, Sigla: "BRC", Descricao: "Aeroporto San Carlos de Bariloche-ARGENTINA"}
		local_SLC := Local{Codigo: 2, Sigla: "SLC", Descricao: "Aeroporto Arturo Merino Benítez-CHILE"}
		local_CDG := Local{Codigo: 3, Sigla: "CDG", Descricao: "Aeroporto Charles de Gaulle-FRANÇA"}
		//local_ORL := Local{Codigo: 4, Sigla: "ORL", Descricao: "Aeroporto Orlando Executive-ESTADOS UNIDOS"}


		caminhos = append(caminhos, Caminho{Origem: local_GRU.Codigo, Destino: local_SLC.Codigo, Distancia: 60})
		caminhos = append(caminhos, Caminho{Origem: local_GRU.Codigo, Destino: local_BRC.Codigo, Distancia: 30})
		caminhos = append(caminhos, Caminho{Origem: local_SLC.Codigo, Destino: local_CDG.Codigo, Distancia: 20})
		caminhos = append(caminhos, Caminho{Origem: local_BRC.Codigo, Destino: local_SLC.Codigo, Distancia: 20})

		fmt.Printf("\nCaminhos: %v\n", caminhos)

		melhorRota := make([]Caminho, 0, 2)

		locaisVisitados := make([]int,0,10)
		locaisNaoVisitados := make([]int,0,10)

		locaisNaoVisitados = append(locaisNaoVisitados, local_GRU.Codigo)
		locaisNaoVisitados = append(locaisNaoVisitados, local_BRC.Codigo)
		locaisNaoVisitados = append(locaisNaoVisitados, local_SLC.Codigo)
		locaisNaoVisitados = append(locaisNaoVisitados, local_CDG.Codigo)

		VerticeDeVisitas :=

		fmt.Printf("\nproximo: %v\n", melhorRota)
	*/
	graph := dtstructs.NewGraph()
	graph.AddEdge("GRU", "BRC", 10)
	graph.AddEdge("BRC", "SCL", 5)
	graph.AddEdge("GRU", "CDG", 75)
	graph.AddEdge("GRU", "SCL", 20)
	graph.AddEdge("GRU", "ORL", 56)
	graph.AddEdge("ORL", "CDG", 5)
	graph.AddEdge("SCL", "ORL", 20)

	fmt.Printf("\nRotas: %v\n", graph)

	fmt.Println(graph.GetMinorPriceRoute("GRU", "CDG"))
	fmt.Println(graph.GetMinorPriceRoute("BRC", "ORL"))

}
