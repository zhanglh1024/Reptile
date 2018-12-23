package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http connect error!")
		return nil, fmt.Errorf("wrong status is not ok :%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
