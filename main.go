package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cli/GetMinorPriceRoute", handler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nRecebi request do Client CLI: \n%v", r)

	keys, ok := r.URL.Query()["name"]

	name := "guest"

	if ok {

		name = keys[0]
	}

	fmt.Fprintf(w, "Hello %s!", name)
}

/*
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
*/
