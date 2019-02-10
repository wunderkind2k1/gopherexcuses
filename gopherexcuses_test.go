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
