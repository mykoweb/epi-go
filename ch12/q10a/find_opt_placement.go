package ch12q10a

import "sort"

type Building struct {
	Population int
	Distance   int
}

type DistanceSorter []Building

// We need to implement sort.Interface by providing the Len, Less, and Swap
// methods
func (a DistanceSorter) Len() int           { return len(a) }
func (a DistanceSorter) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
func (a DistanceSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func FindOptPlacement(buildings []Building) (location int) {
	sort.Sort(DistanceSorter(buildings))

	totalPop := 0
	for _, building := range buildings {
		totalPop += building.Population
	}
	totalPopIsEven := totalPop%2 == 0
	midPop := totalPop / 2

	accumPop := 0
	for i, building := range buildings {
		accumPop += building.Population
		if totalPopIsEven == false {
			if accumPop >= midPop+1 {
				return building.Distance
			}
		} else {
			if accumPop == midPop {
				return (building.Distance + buildings[i+1].Distance) / 2
			} else if accumPop > midPop {
				return building.Distance
			}
		}
	}

	return buildings[len(buildings)-1].Distance
}
