package libs

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Curl 物件參數
type Curl struct {
}

//New 建構式
func (c Curl) New() *Curl {
	return &c
}

//Post 進行HTTP POST
func (c Curl) Post(url string, params []byte) []byte {
	json := bytes.NewBuffer(params)
	resp, err := http.Post(url, "application/json;charset=utf-8", json)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	return body
}

//PostForm 進行HTTP Form POST
func (c Curl) PostForm(url string, params string, header map[string]string) []byte {
	data := strings.NewReader(params)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Panicln(err)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	return body
}
