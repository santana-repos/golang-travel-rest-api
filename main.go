package main

// Implementa o serviço REST definidos na especificação openapi3 'travels-api_openapi3.yaml'

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"travelling-routes/business"
	"travelling-routes/dtstructs"
	"travelling-routes/utils"
)

var Version = "development"
var counter int
var mutex = &sync.Mutex{}
var b business.Business
var filepath = ""
var graph *dtstructs.Graph = nil

const (
	usage = `Version:	%s

Travels REST API deamon utility

usage: %s <csvFile>

example: %s ./input-routes.csv
`
)

type Response struct {
	Cost  float32  `json:"cost"`
	Route []string `json:"route"`
}

type Request struct {
	Origin      string
	Destination string
	Cost        float32
}

// TODO: perform improviments and refactores as soon as possible
func travelsResourceHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Printf("\nEndpoint Hit: GET /travels; Params: %v\n", r.URL.Query())

		parameters := r.URL.Query()
		cost, route, err := b.RetrieveMinorCostRouteFromCSV(graph, strings.ToUpper(parameters["origin"][0]), strings.ToUpper(parameters["destination"][0]))
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		jData, err := json.Marshal(Response{Cost: cost, Route: route})
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)

		log.Println()

	case http.MethodPost:
		//fmt.Fprintf(w, "resource Travels Method POST")
		fmt.Printf("Endpoint Hit: POST /travels; Params: %v; Body: %v\n", r.URL.Query(), r.Body)
		// Create a new record.

		decoder := json.NewDecoder(r.Body)
		var r Request
		err := decoder.Decode(&r)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		log.Println(r)

		graph.AddEdge(strings.ToUpper(r.Origin), strings.ToUpper(r.Destination), float32(r.Cost))

		err = b.UpdateCSVfromGraph(filepath, graph)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))

		log.Println()

	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, Version, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	noSilent := false
	fileData, err := utils.GetFileData(&noSilent)

	if err != nil {
		utils.ExitGracefully(err)
	}

	if _, err := utils.CheckIfValidFile(&fileData.Filepath); err != nil {
		utils.ExitGracefully(err)
	}

	filepath = fileData.Filepath
	b = business.Business{}
	graph, err = b.BuildGraphFromCSV(filepath)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/travels", travelsResourceHandler)

	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := 8080
	log.Printf("\ntravels-api service started at port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))

}
