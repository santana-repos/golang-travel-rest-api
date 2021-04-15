package csv

import (
	"fmt"
	"testing"
)

func TestLoadCSVroutes(t *testing.T) {

	got, err := LoadCSVroutes()

	if err != nil {
		t.Errorf("error: %v", err)

	}

	fmt.Printf("CSV:\n %v", got)
	//fmt.Println(data.Column1 + " " + data.Column2)
}
