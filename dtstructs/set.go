package dtstructs

import "travelling-routes/csv"

// Reference: https://softwareengineering.stackexchange.com/a/273372

// set implementation for small number of items
type routeSet struct {
	slice []csv.RouteData
}

// functions
func (set *routeSet) Add(p csv.RouteData) {
	if !set.Contains(p) {
		set.slice = append(set.slice, p)
	}
}

func (set routeSet) Contains(p csv.RouteData) bool {
	for _, v := range set.slice {
		if v.Equals(p) {
			return true
		}
	}
	return false
}

func (set routeSet) NumElements() int {
	return len(set.slice)
}

func (set routeSet) GetItems() []csv.RouteData {
	return set.slice
}

func NewCSVRouteSet() routeSet {
	return routeSet{(make([]csv.RouteData, 0, 10))}
}
