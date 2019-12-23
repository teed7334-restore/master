package bots

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
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	header["X-Access-Token"] = c.Token
	body := curl.Post(url, []byte(""), header)
	return body
}
