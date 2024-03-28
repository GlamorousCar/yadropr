package main

import (
	"github.com/kljensen/snowball"
	"github.com/kljensen/snowball/english"
	"strings"
)

func stemSentence(sentence string) map[string]bool {
	mapOfWords := make(map[string]bool)

	splitSentence := strings.Split(sentence, " ")
	for _, word := range splitSentence {
		splitWord := strings.Split(word, "'") //обработка случая i'll и подобные

		if len(splitWord) >= 1 {
			stemmedWord, _ := snowball.Stem(splitWord[0], "english", true)
			mapOfWords[stemmedWord] = true
		}
	}
	return mapOfWords
}

func clearSentence(words map[string]bool) []string {
	clearedSentence := []string{}

	for word := range words {
		if !english.IsStopWord(word) {
			clearedSentence = append(clearedSentence, word)
		}

	}

	return clearedSentence
}
