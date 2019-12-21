package bots

type curl interface {
	Post(url string, params []byte) []byte
}
