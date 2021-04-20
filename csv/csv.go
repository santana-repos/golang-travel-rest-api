package csv

import (
	"encoding/csv"
	"os"
	"strconv"
)

type CSVroute struct {
	Origin      string
	Destination string
	Price       float32
}

func LoadCSVroutes(filepath string) ([]CSVroute, error) {

	routes := make([]CSVroute, 0, 20)

	lines, err := ReadCsv(filepath)
	if err != nil {
		panic(err)
		//return nil, err
	}

	// Loop through lines & turn into object
	for _, line := range lines {

		price, err := strconv.ParseFloat(line[2], 32)
		if err != nil {
			return nil, err
		}

		data := CSVroute{
			Origin:      line[0],
			Destination: line[1],
			Price:       float32(price),
		}

		routes = append(routes, data)
	}

	return routes, nil
}

func LoadCSVlines(filepath string) [][]string {
	lines, err := ReadCsv(filepath)

	if err != nil {
		panic(err)
	}

	return lines
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
// reference: https://golangcode.com/how-to-read-a-csv-file-into-a-struct/
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
