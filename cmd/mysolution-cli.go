package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

var Version = "development"
var silentmode *bool

var travel struct {
	origin      string
	destination string
}

const (
	usage = `Version:	%s

Retrieve the minor cost route for travel from [origin] to [destination]

[verbose mode]
usage: %s <csvFile>

example: %s input-routes.csv
	please enter the route: GRU-CDG
best route: GRU - BRC - SCL - ORL - CDG > $40


[silent mode]
usage: %s -s <csvFile>

Options:
`
)

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		exitGracefully(e)
	}
}

func checkIfValidFile(filename string) (bool, error) {
	// Check if file is CSV
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return false, fmt.Errorf("File %s is not CSV", filename)
	}

	// Check if file does exist
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %s does not exist", filename)
	}

	return true, nil
}

func getFileData() (inputFile, error) {
	// Validate arguments
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A filepath argument is required")
	}

	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := flag.Arg(0)

	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("Only comma or semicolon separators are allowed")
	}

	return inputFile{fileLocation, *separator, *pretty}, nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	silentmode = flag.Bool("s", false, "activate silent mode")
	flag.Usage = func() {
		flag.StringVar(&travel.origin, "origin", travel.origin, "origin airport code. ex: GRU")
		flag.StringVar(&travel.destination, "destination", travel.destination, "destination airport code. ex: CDG")
		fmt.Fprintf(flag.CommandLine.Output(), usage, Version, os.Args[0], os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	log.Printf(":\tPassou flag S? %t", *silentmode)

	fileData, err := getFileData()

	if err != nil {
		exitGracefully(err)
	}

	if _, err := checkIfValidFile(fileData.filepath); err != nil {
		exitGracefully(err)
	}
}

/*
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var config struct { // [1]
	port int
	host string
}

const (
	usage = `usage: %s
Run HTTP server

Options:
`
)

func main() {
	flag.IntVar(&config.port, "port", config.port, "port to listen on")    // [2]
	flag.StringVar(&config.host, "host", config.host, "host to listen on") // [3]
	flag.Usage = func() {                                                  // [4]
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse() // [5]

	http.HandleFunc("/", handler)
	addr := fmt.Sprintf("%s:%d", config.host, config.port)
	fmt.Printf("server ready on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}

}

func init() { // [6]
	// Set defaults
	s := os.Getenv("HTTPD_PORT")
	p, err := strconv.Atoi(s)
	if err == nil {
		config.port = p
	} else {
		config.port = 8080
	}

	h := os.Getenv("HTTPD_HOST")
	if len(h) > 0 {
		config.host = h
	} else {
		config.host = "localhost"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gophers\n")
}
*/
