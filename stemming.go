package main

import (
	"github.com/kljensen/snowball"
	"github.com/kljensen/snowball/english"
	"strings"
	"unicode"
)

func stemWords(words []string) []string {
	var stemmedWords []string
	for _, word := range words {
		splitWord := strings.Split(word, "'")
		stemmedWord, _ := snowball.Stem(splitWord[0], "english", true)
		stemmedWords = append(stemmedWords, stemmedWord)

	}
	return stemmedWords
}

func clearSentence(words []string) []string {
	mapOfWords := make(map[string]interface{})

	var clearedSentence []string

	for _, word := range words {
		if !english.IsStopWord(word) {
			mapOfWords[word] = true
		}

	}
	for word := range mapOfWords {
		clearedSentence = append(clearedSentence, word)
	}

	return clearedSentence
}

func Normalize(sentence string) []string {
	f := func(c rune) bool {
		return (unicode.IsPunct(c) || unicode.IsSpace(c)) && c != '\''
	}

	splitSentence := strings.FieldsFunc(sentence, f)

	stemmedSentence := stemWords(splitSentence)

	clearedSentence := clearSentence(stemmedSentence)

	return clearedSentence
}
