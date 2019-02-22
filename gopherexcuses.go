package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const url = "http://developerexcuses.com"

// LoaderStruct is just to bind sth to the Loader interface to allow testing
type LoaderStruct struct{}

// Loader defines methods that imrprove testabilty
type Loader interface {
	getFromURL() (string, error)
	extract(htmlSource string) (string, error)
}

func (loader LoaderStruct) getFromURL() (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body[:]), nil
}

func (loader LoaderStruct) extract(htmlSource string) (string, error) {
	rx := regexp.MustCompile("<a.*?>(.*?)</a>")
	foundLink := rx.FindAllStringSubmatch(htmlSource, -1)
	if foundLink == nil {
		return "", errors.New("Oops no excuse found")
	}
	return foundLink[0][1], nil
}

func main() {
	err := loadExcuse(&LoaderStruct{})
	fmt.Println(handleErrorIfExists("Something went wrong: %s", err))
}

func loadExcuse(loader Loader) error {
	htmlSource, err := loader.getFromURL()
	if err != nil {
		return err
	}
	excuse, err := loader.extract(htmlSource)
	if err != nil {
		return err
	}
	fmt.Println(excuse)
	return nil
}

func handleErrorIfExists(message string, err error) string {
	if err != nil {
		return fmt.Sprintf(message, err)
	}
	return ""
}
