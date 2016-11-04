package ch13q13a

func FindAllStringDecomp(sentence string, words []string) []int {
	// Assume we don't have to check for bad args
	foundDecomp := true
	wordLength := len(words[0])
	result := []int{}

	if len(sentence) < wordLength {
		return result
	}

	wordFreq := make(map[string]int)
	foundWordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	for i := 0; i < wordLength; i++ {
		foundDecomp = true
		result = []int{}
		foundWordFreq = make(map[string]int)

		findStringDecomp(i, sentence[i:], words, &wordFreq, &foundWordFreq, &result)

		for word, _ := range wordFreq {
			if foundWordFreq[word] != wordFreq[word] {
				foundDecomp = false
			}
		}
		if foundDecomp == true {
			break
		}
	}

	if len(result) != len(words) {
		return []int{}
	}

	return result
}

func findStringDecomp(start int, sentence string, words []string, wordFreq, foundWordFreq *map[string]int, result *[]int) {
	wordLength := len(words[0])
	if len(sentence) < wordLength {
		return
	}

	currentWord := sentence[:wordLength]
	if (*wordFreq)[currentWord] == 0 {
		if len(*result) < len(words) { // we found a partial result, but not consecutive
			*result = []int{}
		}
	} else {
		*result = append(*result, start)
		(*foundWordFreq)[currentWord]++
	}

	findStringDecomp(start+wordLength, sentence[wordLength:], words, wordFreq, foundWordFreq, result)
}
