package picker

import (
	"fmt"
	"math"
	"testing"
)

func Test_vowel(t *testing.T) {
	t.Parallel()
	t.Run("check vowel weighting", func(t *testing.T) {
		// Run sum * x times and check for counts within 5% of expected based on weighting
		iterations := 10000
		counts := make(map[rune]int)
		for i := 0; i < sum(vowelWeights)*iterations; i++ {
			counts[vowel()] += 1
		}
		for letter, weight := range vowelWeights {
			difference := math.Abs(float64(counts[letter] - weight*iterations))
			p := difference / float64(weight*iterations) * 100
			if p > 5 {
				t.Errorf("Expected count within 5 percent of weight, got %v", p)
			}

		}
	})
}

func Test_consonant(t *testing.T) {
	t.Parallel()
	t.Run("check consonant weighting", func(t *testing.T) {
		// Run sum * x times and check for counts within 5% of expected based on weighting
		iterations := 10000
		counts := make(map[rune]int)
		for i := 0; i < sum(consonantWeights)*iterations; i++ {
			counts[consonant()] += 1
		}
		for letter, weight := range consonantWeights {
			difference := math.Abs(float64(counts[letter] - weight*iterations))
			p := difference / float64(weight*iterations) * 100
			if p > 5 {
				t.Errorf("Expected count within 5 percent of weight, got %v", p)
			}

		}
	})
}

func TestSelect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		vowels     int
		consonants int
		wantErr    bool
	}{
		{vowels: 0, consonants: 9, wantErr: true},
		{vowels: 1, consonants: 8, wantErr: true},
		{vowels: 2, consonants: 7, wantErr: true},
		{vowels: 3, consonants: 6},
		{vowels: 4, consonants: 5},
		{vowels: 5, consonants: 4},
		{vowels: 6, consonants: 3, wantErr: true},
		{vowels: 7, consonants: 2, wantErr: true},
		{vowels: 8, consonants: 1, wantErr: true},
		{vowels: 9, consonants: 0, wantErr: true},
		{vowels: 100, consonants: 0, wantErr: true},
		{vowels: 0, consonants: 100, wantErr: true},
		{vowels: 100, consonants: 100, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v vowels %v consonants", tt.vowels, tt.consonants), func(t *testing.T) {
			got, err := Select(tt.vowels, tt.consonants)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we want an error dont bother checking the rest
			if tt.wantErr {
				return
			}

			if len(got) != 9 {
				t.Errorf("got %v letters, expected 9", len(got))
			}

			gotVowels, gotCons := 0, 0
			for _, letter := range got {
				if _, found := vowelWeights[letter]; found {
					gotVowels += 1
				}
				if _, found := consonantWeights[letter]; found {
					gotCons += 1
				}
			}
			if gotVowels != tt.vowels {
				t.Errorf("got %v vowels, wanted %v", gotVowels, tt.vowels)
			}
			if gotCons != tt.consonants {
				t.Errorf("got %v consonants, wanted %v", gotCons, tt.consonants)
			}
		})
	}
}
