package app

import (
	"strings"
	"sort"
)

type WordOccurrence struct {
	Word string `json:"word"`
	Cnt  int    `json:"occurrences"`
}

// Sort DESC func
type byCntDesc []WordOccurrence
func (a byCntDesc) Len() int           { return len(a) }
func (a byCntDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCntDesc) Less(i, j int) bool { return a[i].Cnt > a[j].Cnt }

func CountOccurrences(text string) ([]WordOccurrence, error) {
	var err error
	var wordOccurrences []WordOccurrence

	// Remove unwanted chars like numbers and punctuation
	cleanText := textFilter(text)

	// Separate text in to slice of words
	words := strings.Fields(cleanText)

	// Count word occurrences
	countedWords := make(map[string]int)
	for _, word := range words {
		if countedWords[word] != 0 {
			countedWords[word]++
		} else {
			countedWords[word] = 1
		}
	}

	// All all the counted results to the output/result struct
	for word, wordCnt := range countedWords {
		wordOccurrences = append(wordOccurrences, WordOccurrence{
			Word: word,
			Cnt:  wordCnt,
		})
	}

	// Sort DESC by occurrences
	sort.Sort(byCntDesc(wordOccurrences))

	return wordOccurrences, err
}

// Removes most common unwanted chars
// TODO Consider using regex for better filtering of edge cases
func textFilter(t string) string {
	return strings.NewReplacer(
		".", "",
		",", "",
		";", "",
		":", "",
		"1", "",
		"2", "",
		"3", "",
		"4", "",
		"5", "",
		"6", "",
		"7", "",
		"8", "",
		"9", "",
	).Replace(t)
}
