package bots

import (
	"encoding/json"
	"log"
)

//Reserve 預約用資料結構
type Reserve struct {
	Level string
	Token string
}

//New 建構式
func (r Reserve) New(level string, token string) *Reserve {
	r.Level = level
	r.Token = token
	return &r
}

//Reserve 預約寵物
func (r *Reserve) Reserve(url string, curl curl) []byte {
	url = url + "/api/pets/" + r.Level + "/reserve"
	params, err := json.Marshal(r)
	if err != nil {
		log.Panicln(err)
	}
	body := curl.Post(url, params)
	return body
}
