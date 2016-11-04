package ch13q13a

import "testing"

func TestFindStringDecomp(t *testing.T) {
	sentence := "amanaplanacanal"
	words := []string{"can", "apl", "ana"}
	result := FindAllStringDecomp(sentence, words)
	// result should be [4 7 10]
	if len(result) != 3 {
		t.Errorf("Result length expected to be 3 but was %v", len(result))
	}
	if result[0] != 4 {
		t.Errorf("First result val expected to be 4 but was %v", result[0])
	}

	sentence = "aaaaaaaaaaaaaaa"
	words = []string{"aa", "aa"}
	result = FindAllStringDecomp(sentence, words)
	// result should be []
	if len(result) != 0 {
		t.Errorf("Result length expected to be 0 but was %v", len(result))
	}

	sentence = "bbbaaaabb"
	words = []string{"aa", "aa"}
	result = FindAllStringDecomp(sentence, words)
	// result should be [3 5]
	if len(result) != 2 {
		t.Errorf("Result length expected to be 2 but was %v", len(result))
	}
	if result[1] != 5 {
		t.Errorf("Last result val expected to be 5 but was %v", result[1])
	}

	sentence = "bbaabbaa"
	words = []string{"aa", "aa"}
	result = FindAllStringDecomp(sentence, words)
	// result should be []
	if len(result) != 0 {
		t.Errorf("Result length expected to be 0 but was %v", len(result))
	}

	sentence = "aaaabb"
	words = []string{"aa", "aa"}
	result = FindAllStringDecomp(sentence, words)
	// result should be [0 2]
	if len(result) != 2 {
		t.Errorf("Result length expected to be 2 but was %v", len(result))
	}
	if result[0] != 0 {
		t.Errorf("First result val expected to be 0 but was %v", result[0])
	}

	sentence = "bbaaaaaa"
	words = []string{"aaa", "aaa"}
	result = FindAllStringDecomp(sentence, words)
	// result should be [2 5]
	if len(result) != 2 {
		t.Errorf("Result length expected to be 2 but was %v", len(result))
	}
	if result[1] != 5 {
		t.Errorf("Last result val expected to be 5 but was %v", result[1])
	}
}
