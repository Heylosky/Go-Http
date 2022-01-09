package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urls := "http://192.168.3.218/"
	var err error
	var res *http.Response
	datas := url.Values{
		"name": {"TK"},
	}
	req, err := http.NewRequest("GET", urls, strings.NewReader(datas.Encode()))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("This-is-a-test", "666")
	client := &http.Client{}
	res, err = client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	fmt.Println(res.Header)
}
