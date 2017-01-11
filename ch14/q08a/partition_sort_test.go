package ch14q08a

import "testing"

func TestPartitionSort(t *testing.T) {
	students := []Student{
		Student{"E", 2},
		Student{"F", 3},
		Student{"G", 4},
		Student{"H", 2},
		Student{"I", 2},
		Student{"A", 1},
		Student{"B", 2},
		Student{"C", 3},
		Student{"D", 4},
	}

	PartitionSort(students)

	firstIndexForAge1 := index(students, 1)
	if students[firstIndexForAge1].age != 1 {
		t.Errorf("students array was not sorted properly")
	}

	firstIndexForAge2 := index(students, 2)
	for i := 0; i < 4; i++ {
		if students[firstIndexForAge2+i].age != 2 {
			t.Errorf("students array was not sorted properly")
		}
	}

	firstIndexForAge3 := index(students, 3)
	for i := 0; i < 2; i++ {
		if students[firstIndexForAge3+i].age != 3 {
			t.Errorf("students array was not sorted properly")
		}
	}

	firstIndexForAge4 := index(students, 4)
	for i := 0; i < 2; i++ {
		if students[firstIndexForAge4+i].age != 4 {
			t.Errorf("students array was not sorted properly")
		}
	}
}

func index(students []Student, value int) int {
	for i, student := range students {
		if student.age == value {
			return i
		}
	}

	return -1
}
