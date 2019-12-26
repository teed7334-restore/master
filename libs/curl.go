package libs

import (
	"fmt"
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

//GetMockHeader 取得偽裝成Firefox的Header
func (c *Curl) GetMockHeader(domain, prevURL string) map[string]string {
	header := make(map[string]string)
	header["Accept"] = "*/*"
	header["Connection"] = "keep-alive"
	header["Content-Length"] = "57"
	header["Content-Type"] = "application/json; charset=utf-8"
	header["Host"] = domain
	header["Origin"] = fmt.Sprintf("https://%s", domain)
	header["Referer"] = fmt.Sprintf("%s", prevURL)
	header["TE"] = "Trailers"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0"
	header["Accept-Language"] = "zh-TW,en-US;q=0.7,en;q=0.3"
	return header
}

//Post 進行HTTP POST
func (c *Curl) Post(url string, params []byte, header map[string]string) []byte {
	data := &strings.Reader{}
	json := string(params)
	if json != "" {
		data = strings.NewReader(json)
	}
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
