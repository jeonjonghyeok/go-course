package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "log"
	_ "net/http"
)

type someStruct struct {
	testVal int
	SomeInt int `json:"jsonSomeInt" xml:"xmlSomeInt"` //대문자해야 함.그래야 json라이브러리가 읽음
}

func main() {
	str, err := json.Marshal(someStruct{testVal: 1, SomeInt: 2})
	fmt.Println(string(str), err)
	str, err = xml.Marshal(someStruct{testVal: 1, SomeInt: 2})
	fmt.Println(string(str), err)

	var s someStruct 
	if err := json.Unmarshal([]byte(`{"jsonSumint"}`),&s)

}
