package main

import (
	"errors"
	"testing"
)

type ErrorLoader struct {
}

type ErrorMatcher struct {
}

func returnEmptyStringAndAlwaysErr() (string, error) {
	return "", errors.New("Always an error")
}

func (errorLoader ErrorLoader) getFromURL() (string, error) {
	return returnEmptyStringAndAlwaysErr()
}

func (errorLoader ErrorLoader) extract(htmlSource string) (string, error) {
	return htmlSource, nil
}

func (errorLoader ErrorMatcher) getFromURL() (string, error) {
	return "", nil
}

func (errorLoader ErrorMatcher) extract(htmlSource string) (string, error) {
	return returnEmptyStringAndAlwaysErr()
}

func TestLoadExcuseURLError(t *testing.T) {
	err := loadExcuse(&ErrorLoader{})

	if err == nil {
		t.Fatalf("Expect an error when producing errors in http retrieval and htmlsource matching: %s", err)
	}
}

func TestLoadExcuseMatchError(t *testing.T) {
	err := loadExcuse(&ErrorMatcher{})

	if err == nil {
		t.Fatalf("Expect an error when producing errors in http retrieval and htmlsource matching: %s", err)
	}
}

func TestHandleErrorIfExists(t *testing.T) {
	testString := "this is the message: this is an error"

	outString := handleErrorIfExists("this is the message: %s", errors.New("this is an error"))

	if outString != testString {
		t.Fatalf("did not recieve the expected message: %s", outString)
	}
}

func TestSimpleExtraction(t *testing.T) {
	testLoader := &LoaderStruct{}
	link, err := testLoader.extract("<a>http://www.thisisthelink.com</a>")
	if err != nil {
		t.Fatalf("Link not extracted. Error: %s", err)
	}
	if link != "http://www.thisisthelink.com" {
		t.Fatalf("Extracted wrong Link: %s", link)
	}
}

func TestAdvancedExtraction(t *testing.T) {
	testLoader := &LoaderStruct{}
	link, err := testLoader.extract(`<body>div style='font-size:3em'>
	<script><!-- alert('this is js');--></script>
	<a href='http://www.foo.com'>http://www.thisisthelink2.com</a></div><body>`)
	if err != nil {
		t.Fatalf("Link not extracted. Error: %s", err)
	}
	if link != "http://www.thisisthelink2.com" {
		t.Fatalf("Extracted wrong Link: %s", link)
	}
}

func TestNothingfoundExtraction(t *testing.T) {
	testLoader := &LoaderStruct{}
	link, err := testLoader.extract("<div>no link here</div>")
	if link != "" {
		t.Fatalf("No link expected. Found link: %s", link)
	}
	if err == nil {
		t.Fatalf("Expected error not returned")
	}
}
