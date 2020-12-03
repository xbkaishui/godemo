package ch01

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_File_Read(t *testing.T){
	counts := make(map[string] int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n",err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line] ++
		}
		t.Log("")
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
