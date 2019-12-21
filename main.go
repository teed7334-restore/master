package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"

	_ "github.com/joho/godotenv/autoload"
	"github.com/teed7334-restore/master/bots"
	"github.com/teed7334-restore/master/libs"
)

//loginResponse 登入後回傳參數
type loginResponse struct {
	ReturnCode string `json:"returnCode"`
	Info       struct {
		SessionID string   `json:"sessionId"`
		MemberID  string   `json:"memberId"`
		Token     string   `json:"token"`
		Team      []string `json:"team"`
		Expire    string   `json:"expire"`
	}
}

//reserveResponse 預約後回傳參數
type reserveResponse struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	Info          struct {
		Hash      string `json:"hash"`
		Amount    int32  `json:"amount"`
		Op        string `json:"op"`
		IDPet     int32  `json:"idPet"`
		Level     int32  `json:"level"`
		Timestamp string `json:"timestamp"`
		Slot      int32  `json:"slot"`
	}
	StatusCode int32 `json:"statusCode"`
}

//catResponse 抓取後回傳參數
type catResponse struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	StatusCode    int32  `json:"statusCode"`
}

//main 主程式
func main() {
	cpu := runtime.NumCPU()
	thread := 5

	level := os.Getenv("level")
	id := os.Getenv("id")
	lr := doLogin()

	rr := doReserve(level, lr.Info.Token)
	if rr.ReturnCode != "000000" {
		log.Println(rr.ReturnMessage)
	}

	ch := make(chan *catResponse, cpu*thread)
	for i := 0; i < cpu; i++ {
		for j := 0; j < thread; j++ {
			go func() {
				ch <- doCat(level, id, lr.Info.Token)
			}()
		}
	}

	for i := 0; i < cpu*thread; i++ {
		message := <-ch
		log.Println(message.ReturnMessage)
	}
}

//doCat 進行抓取
func doCat(level string, id string, token string) *catResponse {
	url := os.Getenv("url")
	ct := bots.Cat{}.New(id, level, token)
	c := libs.Curl{}.New()
	body := ct.Cat(url, c)
	cr := &catResponse{}
	err := json.Unmarshal(body, cr)
	if err != nil {
		log.Panicln(err)
	}
	return cr
}

//doReserve 進行預約
func doReserve(level string, token string) *reserveResponse {
	url := os.Getenv("url")
	r := bots.Reserve{}.New(level, token)
	c := libs.Curl{}.New()
	body := r.Reserve(url, c)
	rr := &reserveResponse{}
	err := json.Unmarshal(body, rr)
	if err != nil {
		log.Panicln(err)
	}
	return rr
}

//doLogin 進行登入並取得Token
func doLogin() *loginResponse {
	url := os.Getenv("url")
	username := os.Getenv("username")
	password := os.Getenv("password")
	l := bots.Login{}.New(username, password)
	c := libs.Curl{}.New()
	body := l.GetToken(url, c)
	lr := &loginResponse{}
	err := json.Unmarshal(body, lr)
	if err != nil {
		log.Panicln(err)
	}
	return lr
}
