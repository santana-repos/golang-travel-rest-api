package csv

import (
	"encoding/csv"
	"os"
)

type CSVroute struct {
	Origin      string
	Destination string
	Cost        float32
}

func (p1 CSVroute) Equals(p2 CSVroute) bool {
	return (p1.Origin == p2.Origin) && (p1.Destination == p2.Destination) && (p1.Cost == p2.Cost)
}

func LoadCSVlines(filepath string) ([][]string, error) {
	lines, err := ReadCsv(filepath)

	if err != nil {
		return nil, err
	}

	return lines, nil
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
