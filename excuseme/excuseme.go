package excuseme

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const url = "http://developerexcuses.com"

// https://www.useragents.me/
const userAgentChrome = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.3"

// LoaderStruct is just to bind sth to the Loader interface to allow testing
type LoaderStruct struct{}

// Loader defines methods that imrprove testabilty
type Loader interface {
	getFromURL() (string, error)
	extract(htmlSource string) (string, error)
}

func (loader LoaderStruct) getFromURL() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgentChrome)
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Weired: couldn't close body")
			os.Exit(1)
		}
	}(resp.Body)
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

// LoadExcuse loads the excuse with the given Loader - currentyl using the site developerexcuses.com
func LoadExcuse(loader Loader) error {
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

// HandleErrorIfExists returns the message if the the err != nil else ""
func HandleErrorIfExists(message string, err error) string {
	if err != nil {
		return fmt.Sprintf(message, err)
	}
	return ""
}
