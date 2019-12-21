package libs

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
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
