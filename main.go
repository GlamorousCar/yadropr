package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	inputString := flag.String("s", "", "Предложение для стемминга")

	flag.Parse()

	stemmedWords := stemSentence(*inputString)
	clearedSentence := clearSentence(stemmedWords)

	fmt.Println(strings.Join(clearedSentence, " "))

}
