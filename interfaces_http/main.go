package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}


func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	l := logWriter{}
	io.Copy(l, resp.Body)


}

func (logWriter) Write(bs []byte) (int, error) {
	// fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}