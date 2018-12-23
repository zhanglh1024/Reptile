package parser

import (
	"Reptile/regex/engine"
	"fmt"
	"regexp"
)

const CityList  = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a> `

func ParseCityList(content []byte)  {
	re := regexp.MustCompile(CityList)
	city := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m:= range city{
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParse,

		})
		fmt.Printf("City:%s url:%s all:%s\n", m[2], m[1], m)
	}
}
