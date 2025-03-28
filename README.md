# gopherexcuses

This app is a cli to download excuses for you golang developers from developerexcuses.com

# How to use

Go get it.

```shell
go install github.com/wunderkind2k1/gopherexcuses@latest
```
and then call


```shell
gopherexcuses
```

# Have fun with it like this:

```shell
# dont be greedy and download to many. repect the service.

for i in {1..10}; do gopherexcuses >> excuses.txt; done
```


# You can embedd the "service" in your own code by importing it:

> import "github.com/wunderkind2k1/gopherexcuses/excuseme"


## Example

```Go
package main

import (
	"fmt"
	"github.com/wunderkind2k1/gopherexcuses/excuseme"
)

func main() {
	err := excuseme.LoadExcuse(&excuseme.LoaderStruct{})
	fmt.Println(excuseme.HandleErrorIfExists("Something went wrong: %s", err))
}


```
