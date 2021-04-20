package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

type RouteData struct {
	Origin      string
	Destination string
	Cost        float32
}

func (p1 RouteData) Equals(p2 RouteData) bool {
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

// reference: https://golang.org/pkg/encoding/csv/#example_Writer
func WriteCsv(filename string, routeData []RouteData, fakeWrite bool) error {

	var writer *csv.Writer
	if fakeWrite {
		writer = csv.NewWriter(os.Stdout)
	} else {
		// Write CSV file
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		writer = csv.NewWriter(file)
	}

	for _, route := range routeData {
		record := []string{route.Origin, route.Destination, fmt.Sprintf("%.2f", route.Cost)}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to csv: %v", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	writer.Flush()

	if err := writer.Error(); err != nil {
		return fmt.Errorf("unnexpected error: %v", err)
	}

	return nil
}
