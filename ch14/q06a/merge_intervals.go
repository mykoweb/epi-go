package ch14q06a

type Interval [2]int

func mergeIntervals(sortedIntervals []Interval, newInterval Interval) []Interval {
	var result []Interval

	index := 0

	for index < len(sortedIntervals) && sortedIntervals[index][1] < newInterval[0] {
		result = append(result, sortedIntervals[index])
		index++
	}

	for index < len(sortedIntervals) && sortedIntervals[index][0] <= newInterval[1] {
		newInterval[0] = min(sortedIntervals[index][0], newInterval[0])
		newInterval[1] = max(sortedIntervals[index][1], newInterval[1])
		index++
	}

	result = append(result, newInterval)

	if index < len(sortedIntervals) {
		result = append(result, sortedIntervals[index:]...)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
