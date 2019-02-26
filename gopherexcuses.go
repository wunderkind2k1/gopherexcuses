package main

import (
	"fmt"
	"github.com/wunderkind2k1/gopherexcuses/excuseme"
)

func main() {
	err := excuseme.LoadExcuse(&excuseme.LoaderStruct{})
	fmt.Println(excuseme.HandleErrorIfExists("Something went wrong: %s", err))
}
