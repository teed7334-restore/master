package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

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

//notifyResponse 寄送訊息回傳參數
type notifyResponse struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

//main 主程式
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	level := os.Args[1]
	id := os.Args[2]
	count, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Println(err)
	}

	lr := doLogin()

	content := ""

	rr := doReserve(level, lr.Info.Token)
	if rr.ReturnCode != "000000" {
		content = fmt.Sprintf("[%s] %s", rr.ReturnCode, rr.ReturnMessage)
		log.Println(content)
	}

	ch := make(chan *catResponse, count)
	for i := 0; i < count; i++ {
		runtime.Gosched()
		go func() {
			ch <- doCat(level, id, lr.Info.Token)
		}()
	}

	message := &catResponse{}
	logs := fmt.Sprintf("一共發出 %d 顆寶貝球\n", count)
	success := 0
	failure := 0
	for i := 0; i < count; i++ {
		message = <-ch
		content = fmt.Sprintf("[%s] %s", message.ReturnCode, message.ReturnMessage)
		log.Println(content)
		if message.ReturnCode == "000000" {
			success++
		} else {
			failure++
		}
	}
	logs += fmt.Sprintf("有抓到一共 %d 隻\n", success)
	logs += fmt.Sprintf("沒抓到一共 %d 隻", failure)
	doSend(logs)
}

func doSend(message string) *notifyResponse {
	notify := os.Getenv("notify")
	token := os.Getenv("token")
	n := bots.Notify{}.New(notify, token)
	c := libs.Curl{}.New()
	body := n.Send(message, c)
	nr := &notifyResponse{}
	err := json.Unmarshal(body, nr)
	if err != nil {
		log.Panicln(err)
	}
	return nr
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
