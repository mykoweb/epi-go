package ch14q08a

type Student struct {
	name string
	age  int
}

type offsetTuple struct {
	startOffset int
	endOffset   int
}

// PartitionSort() groups all students with the same age together. The order of
// appearance of each resulting age group does not matter.
//
// The sorting will be done in-place.
func PartitionSort(students []Student) {
	ageCount := make(map[int]int)
	ageOffset := make(map[int]*offsetTuple)
	overallOffset := 0

	// Create a map that counts the number of students in each age group.
	for _, student := range students {
		ageCount[student.age]++
	}

	// Create a map that stores the start and stop offsets for each age group
	// within the students slice.
	for age, ageNum := range ageCount {
		ageOffset[age] = &offsetTuple{startOffset: overallOffset, endOffset: overallOffset + ageNum}
		overallOffset += ageNum
	}

	// Perform in-place swapping of student objects until all the students with
	// the same age are grouped together.
	for len(ageOffset) > 0 {
		for age, offset := range ageOffset {
			if offset.startOffset == offset.endOffset {
				delete(ageOffset, age)
			} else {
				studentAgeAtOffset := students[offset.startOffset].age
				if studentAgeAtOffset != age {
					students[offset.startOffset], students[ageOffset[studentAgeAtOffset].startOffset] = students[ageOffset[studentAgeAtOffset].startOffset], students[offset.startOffset]
					ageOffset[studentAgeAtOffset].startOffset++
				} else {
					offset.startOffset++
				}
			}
		}
	}
}
