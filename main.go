package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	inputString := flag.String("s", "", "Предложение для стемминга")

	flag.Parse()

	processedSentence := Normalize(*inputString)

	fmt.Println(strings.Join(processedSentence, " "))
}
