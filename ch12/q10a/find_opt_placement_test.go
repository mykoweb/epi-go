package ch12q10a

import "testing"

func TestFindOptPlacement(t *testing.T) {
	buildings := []Building{}

	// Total pop is odd
	buildings = append(buildings, Building{101, 10})
	buildings = append(buildings, Building{100, 2})
	buildings = append(buildings, Building{100, 20})
	buildings = append(buildings, Building{100, 100})

	location := FindOptPlacement(buildings)

	if location != 10 {
		t.Errorf("Expected distance of 10 but got %v", location)
	}

	// Total pop is even
	buildings = []Building{}
	buildings = append(buildings, Building{100, 1000})
	buildings = append(buildings, Building{100, 1})
	buildings = append(buildings, Building{100, 2000})
	buildings = append(buildings, Building{100, 3500})

	location = FindOptPlacement(buildings)

	if location != 1500 {
		t.Errorf("Expected distance of 1500 but got %v", location)
	}

	// Location should be at last building
	buildings = []Building{}
	buildings = append(buildings, Building{1, 1000})
	buildings = append(buildings, Building{1, 1})
	buildings = append(buildings, Building{1000, 3500})
	buildings = append(buildings, Building{1, 2000})

	location = FindOptPlacement(buildings)

	if location != 3500 {
		t.Errorf("Expected distance of 3500 but got %v", location)
	}
}
