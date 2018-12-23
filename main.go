 package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http connect error!")
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	printAllCity(all)
}

 func printAllCity(content []byte) {
	 re := regexp.MustCompile(`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a> `)
	 city := re.FindAllSubmatch(content, -1)

	 for _, m:= range city{
	 		fmt.Printf("City:%s url:%s all:%s\n", m[2], m[1], m)
	 }

 }
