package phone_number

import (
	"fmt"
	"testing"
)

func TestGenerateWords(t *testing.T) {
	result := GenerateWords("23")
	fmt.Println(result) // [ad ae af bd be bf cd ce cf]
	if len(result) != 9 {
		t.Errorf("Result length was incorrect")
	}

	if result[0] != "ad" {
		t.Errorf("Expected 'ad' but got", result[0])
	}

	if result[len(result)-1] != "cf" {
		t.Errorf("Expected 'cf' but got", result[len(result)-1])
	}
}
