package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const url = "http://developerexcuses.com"

func main() {

	loadExcuse()
}

func loadExcuse() error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	htmlSource := string(body[:])
	rx := regexp.MustCompile("<a.*?>(.*?)</a>")
	foundLink := rx.FindAllStringSubmatch(htmlSource, -1)
	if foundLink == nil {
		fmt.Println("Opps no excuse found")
	}
	fmt.Println(foundLink[0][1])
	return nil
}

func handleErrorIfExists(message string, err error) {
	if err != nil {
		fmt.Printf(message, err)
	}
}
