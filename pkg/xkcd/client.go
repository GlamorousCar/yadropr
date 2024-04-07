package xkcd

import (
	"encoding/json"
	"fmt"
	"github.com/GlamorousCar/yadropr/pkg/words"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Comics struct {
	Num        int    `json:"num"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}

func GetComicsFromSite(url string) (Comics, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return Comics{}, 400, err
	}
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return Comics{}, 400, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	var comics Comics
	json.Unmarshal(body, &comics)

	if err != nil {
		fmt.Printf("error %s", err)
		return Comics{}, 400, err
	}

	return comics, resp.StatusCode, nil
}

func GetIDLastComics(url string) int {
	comics, _, err := GetComicsFromSite(url + "/info.0.json")

	if err != nil {
		fmt.Printf("error %s", err)
		return 0
	}
	return comics.Num

}

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetComics(url string, maxRecord int, showOutput bool) []Comics {
	var countComics int
	if maxRecord == 0 {
		countComics = GetIDLastComics(url)
	} else {
		countComics = maxRecord
	}
	var c []Comics
	var wg sync.WaitGroup
	for i := 1; i < countComics; i++ {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			status, content := doReq(url)
			if status != 404 {
				comics, err := getComicsFromData(content)
				if err != nil {
					return
				}
				c = append(c, comics)

				if showOutput {
					listOfWords := words.Normalize(comics.Transcript + comics.Alt)
					fmt.Printf("{ID = %s\n\turl = %s\n\tkeywords = [%s]\n", strconv.Itoa(comics.Num), comics.Img, strings.Join(listOfWords, ", "))
				}
			} else {
				if showOutput {
					fmt.Println("Обработка исключения", url)
				}
			}
		}(url + "/" + strconv.Itoa(i) + "/info.0.json")
	}
	wg.Wait()

	return c
}

func doReq(url string) (status int, content []byte) {

	resp, err := http.Get(url)

	if err != nil {

		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Println(err)
		return
	}

	return resp.StatusCode, body
}

func getComicsFromData(body []byte) (Comics, error) {
	var comics Comics
	err := json.Unmarshal(body, &comics)
	if err != nil {
		return Comics{}, err
	}

	return comics, nil
}
