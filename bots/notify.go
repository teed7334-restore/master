package bots

import (
	"fmt"
	"net/http"
)

//Notify 物件參數
type Notify struct {
	URL   string
	Token string
}

//New 建構式
func (n Notify) New(url, token string) *Notify {
	n.URL = url
	n.Token = token
	return &n
}

//Send 傳送訊息
func (n *Notify) Send(message string, curl curl) []byte {
	var r http.Request
	r.ParseForm()
	r.Form.Add("message", message)
	params := r.Form.Encode()
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", n.Token)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	body := curl.Post(n.URL, []byte(params), header)
	return body
}
