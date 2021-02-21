package main

import (
	"fmt"
)

type someStruct struct {
	someInt int
}

type student struct {
}

func (s someStruct) String() string {
	return fmt.Sprintf("someint: %d", s.someInt)
}

var s interface{ String() string }

func prettyPrint(s interface{ String() string }) {
	fmt.Println(s.String())
}

func main() {
	s := someStruct{
		someInt: 1,
	}

	/*	resp, err := http.Get("https://www.google.com")
		if err != nil {
			log.Fatal(err)
		}
	*/

	prettyPrint(s)

}
