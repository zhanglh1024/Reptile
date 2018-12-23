package main

import (
	"fmt"
	"regexp"
)

const test  = `my email is zhanglonghu@gmail.com
	my second email is hello@gmail.com.cn
	lllll redis@gmail.com
`

func main(){
	re:= regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(test, -1)
	for _, valu := range match{
		fmt.Println(valu)
	}
	fmt.Println(match)
}
