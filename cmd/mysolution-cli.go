package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"travelling-routes/business"
)

/*
	references:
		https://github.com/Andrew4d3/go-csv2json/blob/d56cb4088a54dcfd21325d5603ce0fceaf1cff5b/csv2json.go#L80
		https://blog.gopheracademy.com/advent-2019/flags/
*/

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
usage: %s -s -origin <origin> -destination <destination> <csvFile>

Options:
`
)

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func checkIfValidFile(filename string) (bool, error) {
	// Check if file is CSV
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return false, fmt.Errorf("file %s is not csv", filename)
	}

	// Check if file does exist
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", filename)
	}

	return true, nil
}

func getFileData(silentmode *bool) (inputFile, error) {
	erromsg := "A filepath argument is required"
	// Validate arguments
	if *silentmode {
		if len(os.Args) < 7 {
			return inputFile{}, errors.New(erromsg)
		}
	} else {
		if len(os.Args) < 2 {
			return inputFile{}, errors.New(erromsg)
		}
	}

	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := os.Args[(len(os.Args) - 1)]
	fmt.Println(fileLocation)

	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("only comma or semicolon separators are allowed")
	}

	return inputFile{fileLocation, *separator, *pretty}, nil
}

func main() {
	silentmode = flag.Bool("s", false, "activate silent mode")
	flag.StringVar(&travel.origin, "origin", "---", "origin airport code. ex: GRU")
	flag.StringVar(&travel.destination, "destination", "---", "destination airport code. ex: CDG")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, Version, os.Args[0], os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Println(len(os.Args))

	fileData, err := getFileData(silentmode)

	if err != nil {
		exitGracefully(err)
	}

	if _, err := checkIfValidFile(fileData.filepath); err != nil {
		exitGracefully(err)
	}

	var parameters string
	if !*silentmode {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nplease enter the route: ")
		parameters, _ = reader.ReadString('\n')
		separator := "-"
		if (len(parameters) != 8) || (string(parameters[3]) != separator) {
			exitGracefully(fmt.Errorf("invalid input parametes: %s", parameters))
		}
		splits := strings.Split(parameters, separator)
		travel.origin = splits[0]
		travel.destination = splits[1]
	}

	/*
		log.Printf("\nPassou flag S? %t", *silentmode)
		log.Printf("\nfilepath: %s", fileData.filepath)
		log.Printf("\nstrings.ToUpper(travel.origin): %s", strings.ToUpper(travel.origin))
		log.Printf("\nstrings.ToUpper(travel.destination): %s", strings.ToUpper(travel.destination))
		log.Printf("\nParameters: %s", parameters)
		log.Printf("\ntravel.origin: %s; travel.destination: %s", travel.origin, travel.destination)
	*/

	b := business.Business{}
	cost, route, err := b.RetrieveMinorCostRouteFromCSV(fileData.filepath, strings.ToUpper(travel.origin), strings.ToUpper(strings.TrimSpace(travel.destination)))
	if err != nil {
		exitGracefully(err)
	}

	sRoute := strings.ReplaceAll(fmt.Sprintf("%v", route), " ", " - ")[1:]
	fmt.Printf("\nbest route: %v > $%2.f\n", strings.TrimSuffix(sRoute, "]"), cost)

	log.Println()
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
