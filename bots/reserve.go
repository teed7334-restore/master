package bots

import (
	"fmt"
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
	api := fmt.Sprintf("%s/api/pets/%s/reserve", url, r.Level)
	prev := fmt.Sprintf("%s/index.html", url)
	header := curl.GetMockHeader(url, prev)
	header["X-Access-Token"] = r.Token
	body := curl.Post(api, []byte(""), header)
	return body
}
