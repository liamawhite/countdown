package picker

import (
	"fmt"
	"math/rand"
	"time"
)

// http://www.thecountdownpage.com/letters.htm
func Select(vowels, consonants int) ([]rune, error) {
	if vowels < 3 {
		return nil, fmt.Errorf("must have at least 3 vowels")
	}
	if consonants < 4 {
		return nil, fmt.Errorf("must have at least 4 consonants")
	}
	if consonants+vowels != 9 {
		return nil, fmt.Errorf("must have 9 letters")
	}

	res := make([]rune, 0, 9)
	for i := 0; i < vowels; i++ {
		res = append(res, vowel())
	}
	for i := 0; i < consonants; i++ {
		res = append(res, consonant())
	}
	
	// Shuffle result
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })

	return res, nil
}

func vowel() rune {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, sum(vowelWeights)
	num := rand.Intn(max-min+1) + min

	for letter, weight := range vowelWeights {
		if num <= weight {
			return letter
		}
		num -= weight
	}
	panic("unable to select vowel, this should not be possible")
}

func consonant() rune {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, sum(consonantWeights)
	num := rand.Intn(max-min+1) + min

	for letter, weight := range consonantWeights {
		if num <= weight {
			return letter
		}
		num -= weight
	}
	panic("unable to select consonant, this should not be possible")
}

func sum(list map[rune]int) int {
	total := 0
	for _, weight := range list {
		total += weight
	}
	return total
}

var vowelWeights map[rune]int = map[rune]int{
	'A': 15,
	'E': 21,
	'I': 13,
	'O': 13,
	'U': 5,
}

var consonantWeights map[rune]int = map[rune]int{
	'B': 2,
	'C': 3,
	'D': 6,
	'F': 2,
	'G': 3,
	'H': 2,
	'J': 1,
	'K': 1,
	'L': 5,
	'M': 4,
	'N': 8,
	'P': 4,
	'Q': 1,
	'R': 9,
	'S': 9,
	'T': 9,
	'V': 1,
	'W': 1,
	'X': 1,
	'Y': 1,
	'Z': 1,
}
