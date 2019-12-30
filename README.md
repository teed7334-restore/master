# master
自動抓取寵物機器人

## 資料夾結構
bots 機器人物件

libs 共用函式庫

.env.swp 置換用設定檔

main.go 主程式

go.mod Golang Module

## 設定教學
1. 置換你的.env.swp成.env
2. 依.env檔裡面要求的，分別輸入你抓寵物用的帳號、密碼、網址(HTTPS://[Domain Name]即可)、Line通知的Token

## 運行教學
預約寵物
```
./master reserve [寵物等級]
```

抓取寵物
```
./master cat [寵物等級] [寵物ID] [進程數]
```