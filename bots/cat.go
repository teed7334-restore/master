package bots

import (
	"fmt"
)

//Cat 抓取用資料結構
type Cat struct {
	ID    string
	Level string
	Token string
}

//New 建構式
func (c Cat) New(id, level, token string) *Cat {
	c.ID = id
	c.Level = level
	c.Token = token
	return &c
}

//Cat 抓取寵物
func (c *Cat) Cat(url string, curl curl) []byte {
	api := fmt.Sprintf("%s/api/pets/%s/panic/%s", url, c.Level, c.ID)
	prev := fmt.Sprintf("%s/index.html", url)
	header := curl.GetMockHeader(url, prev)
	header["X-Access-Token"] = c.Token
	body := curl.Post(api, []byte(""), header)
	return body
}
