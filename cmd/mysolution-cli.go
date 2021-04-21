package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"travelling-routes/business"
	"travelling-routes/utils"
)

/*
	references:
		https://github.com/Andrew4d3/go-csv2json/blob/d56cb4088a54dcfd21325d5603ce0fceaf1cff5b/csv2json.go#L80
		https://blog.gopheracademy.com/advent-2019/flags/
*/

var travel struct {
	origin      string
	destination string
}

var Version = "development"
var silentmode *bool

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

func main() {
	silentmode = flag.Bool("s", false, "activate silent mode")
	flag.StringVar(&travel.origin, "origin", "---", "origin airport code. ex: GRU")
	flag.StringVar(&travel.destination, "destination", "---", "destination airport code. ex: CDG")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, Version, os.Args[0], os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	//fmt.Println(len(os.Args))

	fileData, err := utils.GetFileData(silentmode)

	if err != nil {
		utils.ExitGracefully(err)
	}

	if _, err := utils.CheckIfValidFile(&fileData.Filepath); err != nil {
		utils.ExitGracefully(err)
	}

	var parameters string
	if !*silentmode {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nplease enter the route: ")
		parameters, _ = reader.ReadString('\n')
		separator := "-"
		if (len(parameters) != 8) || (string(parameters[3]) != separator) {
			utils.ExitGracefully(fmt.Errorf("invalid input parametes: %s", parameters))
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
	b = business.Business{}
	graph, err := b.BuildGraphFromCSV(fileData.Filepath)
	if err != nil {
		utils.ExitGracefully(err)
	}
	cost, route, err := b.RetrieveMinorCostRouteFromCSV(graph, strings.ToUpper(travel.origin), strings.ToUpper(strings.TrimSpace(travel.destination)))
	if err != nil {
		utils.ExitGracefully(err)
	}

	sRoute := strings.ReplaceAll(fmt.Sprintf("%v", route), " ", " - ")[1:]
	fmt.Printf("\nbest route: %v > $%2.f\n", strings.TrimSuffix(sRoute, "]"), cost)

	//log.Println()
}
