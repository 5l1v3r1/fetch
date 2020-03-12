// A Command Line Utility To Fetch Websites source code.
// Created By r2dr0dn
// 2020-03-12
package main
import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	if len(os.Args[1:]) !=1 {
		fmt.Printf("Usage: go run fetch.go <Url To Fetch>\n")
	}
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
