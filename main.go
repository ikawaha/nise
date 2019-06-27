package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func usage() string {
	return "nise is a translator from Japanese to Chinese-like fake (hanamogera) string. see github.com/ikawaha/nise"
}

func run(args []string) int {
	var r io.ReadCloser
	if len(args) >= 1 {
		r = ioutil.NopCloser(strings.NewReader(args[0]))
	} else {
		r = os.Stdin
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		fmt.Fprintln(os.Stdout, Filter(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("%v", err)
		return 1
	}
	return 0
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Fprintln(os.Stdout, usage())
			os.Exit(1)
		}
	}
	os.Exit(run(os.Args[1:]))
}
