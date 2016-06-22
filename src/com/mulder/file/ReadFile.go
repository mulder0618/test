package main

import "fmt"
import (
	"os"
)

func main() {
	//写文件
	var fileName string
	fileName = "test.txt"
	fileOut,err := os.Create(fileName)
	defer fileOut.Close()
	if err != nil {
		fmt.Println(fileName,err)
		return
	}
	fileOut.WriteString("this is test")

	//读文件
	var readFileName = "test.txt"
	fileIn,err := os.Open(readFileName)
	defer fileIn.Close()
	if err != nil {
		fmt.Println(readFileName,err)
		return
	}
	buf := make([]byte,1024,1024)
	for  {
		n,_ := fileIn.Read(buf)
		if n==0 {break}
		//os.Stdout.Write(buf)
		os.Stdout.Write(buf[:n])
	}

}