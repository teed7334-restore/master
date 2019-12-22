package bots

import (
	"encoding/json"
	"log"
)

//Cat 抓取用資料結構
type Cat struct {
	ID    string `json:"_"`
	Level string `json:"_"`
	Token string `json:"X-Access-Token"`
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
	url = url + "/api/pets/" + c.Level + "/panic/" + c.ID
	params, err := json.Marshal(c)
	if err != nil {
		log.Panicln(err)
	}
	body := curl.Post(url, params)
	return body
}
