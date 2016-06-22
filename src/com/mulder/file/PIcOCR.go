package main


import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	// This is the simlest way :)
	out := gosseract.Must(gosseract.Params{
		Src:       "/Users/mulder/Downloads/authImg.jpeg",
		Languages: "eng+heb",
	})
	fmt.Println(out)

	// Using client
	client, _ := gosseract.NewClient()
	out, _ = client.Src("/Users/mulder/Downloads/authImg.jpeg").Out()
	fmt.Println(out)
}