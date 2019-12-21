package bots

import (
	"encoding/json"
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
	url = url + "/api/sessions"
	params, err := json.Marshal(l)
	if err != nil {
		log.Panicln(err)
	}
	body := curl.Post(url, params)
	return body
}
