package ch01

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_HTTP_Get(t *testing.T){
	for _, url := range os.Args[4:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error :%v", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v \n", url, err)
			continue
		}
		fmt.Println(string(b))
	}
}
