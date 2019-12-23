package bots

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
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	header["X-Access-Token"] = r.Token
	body := curl.Post(url, []byte(""), header)
	return body
}
