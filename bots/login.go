package bots

import (
	"encoding/json"
	"fmt"
	"log"
)

//Login 登入用資料結構
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//New 建構式
func (l Login) New(username, password string) *Login {
	l.Username = username
	l.Password = password
	return &l
}

//GetToken 取得驗証碼
func (l *Login) GetToken(url string, curl curl) []byte {
	api := fmt.Sprintf("%s/api/sessions", url)
	prev := fmt.Sprintf("%s/login.html", url)
	params, err := json.Marshal(l)
	if err != nil {
		log.Panicln(err)
	}
	header := curl.GetMockHeader(url, prev)
	body := curl.Post(api, params, header)
	return body
}
