package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _,url := range os.Args[1:] {
		prefix := strings.HasPrefix(url, "http://")
		if !prefix {
			url = fmt.Sprintf("http://%s",url)
			fmt.Printf("%s\n",url)
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr,"frtch %v\n",err)
			os.Exit(1)
		}

		fmt.Printf("%v\n",resp.Status)

		bytes, e := io.Copy(os.Stderr,resp.Body)
		resp.Body.Close()
		if e != nil {
			fmt.Fprintf(os.Stderr,"fetch readAll :%s %v\n",url,e)
			os.Exit(1)
		}
		fmt.Printf("%s",bytes)
	}
}

