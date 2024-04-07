package database

import (
	"encoding/json"

	"github.com/GlamorousCar/yadropr/pkg/words"
	"github.com/GlamorousCar/yadropr/pkg/xkcd"
	"io/ioutil"
	"log"
	"strconv"
)

type ComicsContent struct {
	URL      string   `json:"url"`
	Keywords []string `json:"keywords"`
}

func WriteComicsToJson(filename string, comics []xkcd.Comics) {
	data := map[string]ComicsContent{}
	for _, val := range comics {
		listOfWords := words.Normalize(val.Transcript + val.Alt)
		data[strconv.Itoa(val.Num)] = ComicsContent{URL: val.Img, Keywords: listOfWords}

	}
	rawDataOut, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(filename, rawDataOut, 0777)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}
