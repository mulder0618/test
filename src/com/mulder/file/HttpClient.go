package main

import "fmt"
import (
	"net/http"
	"io/ioutil"
	"github.com/opesun/goquery"
)
//import "io/ioutil"

func main() {
	response,_ := http.Get("http://www.baidu.com")
	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
   	bodyContent := string(body)
	//fmt.Println(bodyContent)
	node,_:=goquery.ParseString(bodyContent)
	fmt.Println(node.Find("div").Html())
	fmt.Println("test:",node.Find("#head").Html())
}