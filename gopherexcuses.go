package main

import (
	"./excuseme"
	"fmt"
)

func main() {
	err := excuseme.LoadExcuse(&excuseme.LoaderStruct{})
	fmt.Println(excuseme.HandleErrorIfExists("Something went wrong: %s", err))
}
