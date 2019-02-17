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
