
# Based on https://github.com/mathdroid/covid-19-api

* ### How to add 

  `$ go get https://github.com/izzatzr/covid19api-go`

* ### Example code

```package main

import (
	"fmt"

	covid19 "github.com/izzatzr/covid19api-go"
)

func main() {
	myapp := covid19.New()
	resp := myapp.ByCountrySummary("usa")
	fmt.Println(resp)
}
```

# TODO
* add json prettier as optional return
