package main

import (
	"flag"
	"github.com/GlamorousCar/yadropr/pkg/database"
	"github.com/GlamorousCar/yadropr/pkg/xkcd"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	URL     string `yaml:"source_url"`
	DB_file string `yaml:"db_file"`
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	var appCongig Config

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &appCongig)
	if err != nil {
		panic(err)
	}

	var maxRecord int

	var showOutput bool

	flag.IntVar(&maxRecord, "n", 0, "Вывод на экран максимум n записей")
	flag.BoolVar(&showOutput, "o", false, "Отобразить вывод")

	flag.Parse()

	comics := xkcd.GetComics(appCongig.URL, maxRecord, showOutput)

	database.WriteComicsToJson(appCongig.DB_file, comics)

}
