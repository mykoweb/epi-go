package ch13q13a

import "fmt"

func FindStringDecompBad(sentence string, words []string) []int {
	// Assume for now that we don't need to check for bad inputs
	var foundWordFreq map[string]int
	var result []int
	var didNotFindDecomp bool
	var currSentence, currWord string

	wordFreq := make(map[string]int)

	// sentLength := len(sentence)
	wordLength := len(words[0])

	for _, word := range words {
		wordFreq[word]++
	}

	fmt.Println("wordFreq is", wordFreq)
	for i := 0; i < wordLength; i++ {
		foundWordFreq = make(map[string]int)
		didNotFindDecomp = false

		currSentence = sentence[i:]

		for j := 0; j < len(currSentence); j += wordLength {
			if len(currSentence)-j < wordLength {
				break
			} else {
				fmt.Println("currSentence, j, and j+wordLength and result", currSentence, j, j+wordLength, currSentence[j:j+wordLength])
				currWord = currSentence[j : j+wordLength]
			}
			if wordFreq[currWord] > 0 {
				foundWordFreq[currWord]++
				result = append(result, j+i)
			}
			fmt.Println("foundWordFreq is", foundWordFreq)
			for k, _ := range wordFreq {
				if wordFreq[k] != foundWordFreq[k] {
					didNotFindDecomp = true
					break
				}
			}
			if didNotFindDecomp == false {
				return result
			}
		}
		result = []int{}
	}

	return result
}
