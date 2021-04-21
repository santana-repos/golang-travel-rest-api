package utils

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type InputFile struct {
	Filepath  string
	separator string
	pretty    bool
}

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func CheckIfValidFile(filename *string) (bool, error) {
	// Check if file is CSV
	if fileExtension := filepath.Ext(*filename); fileExtension != ".csv" {
		return false, fmt.Errorf("file %s is not csv", *filename)
	}

	// Check if file does exist
	if _, err := os.Stat(*filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", *filename)
	}

	return true, nil
}

func GetFileData(silentmode *bool) (InputFile, error) {
	erromsg := "A filepath argument is required"
	// Validate arguments
	if *silentmode {
		if len(os.Args) < 7 {
			return InputFile{}, errors.New(erromsg)
		}
	} else {
		if len(os.Args) < 2 {
			return InputFile{}, errors.New(erromsg)
		}
	}

	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := os.Args[(len(os.Args) - 1)]
	//fmt.Println(fileLocation)

	if !(*separator == "comma" || *separator == "semicolon") {
		return InputFile{}, errors.New("only comma or semicolon separators are allowed")
	}

	return InputFile{fileLocation, *separator, *pretty}, nil
}
