package dtstructs

import "travelling-routes/csv"

// Reference: https://softwareengineering.stackexchange.com/a/273372

// set implementation for small number of items
type CSVRouteSet struct {
	slice []csv.CSVroute
}

// functions
func (set *CSVRouteSet) Add(p csv.CSVroute) {
	if !set.Contains(p) {
		set.slice = append(set.slice, p)
	}
}

func (set CSVRouteSet) Contains(p csv.CSVroute) bool {
	for _, v := range set.slice {
		if v.Equals(p) {
			return true
		}
	}
	return false
}

func (set CSVRouteSet) NumElements() int {
	return len(set.slice)
}

func NewCSVRouteSet() CSVRouteSet {
	return CSVRouteSet{(make([]csv.CSVroute, 0, 10))}
}
