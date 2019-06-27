package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	var r io.ReadCloser
	if len(os.Args) >= 2 {
		r = ioutil.NopCloser(strings.NewReader(os.Args[2]))
	}else{
		r = os.Stdin
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)
	for i:=0; scanner.Scan();i++ {
		os.Stdout.WriteString(Filter(scanner.Text()))
		os.Stdout.WriteString("\n")
	}
	if err :=scanner.Err(); err != nil{
		fmt.Errorf("%v", err)
	}
	return
}