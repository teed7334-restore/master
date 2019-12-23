package bots

type curl interface {
	Post(url string, params []byte, header map[string]string) []byte
}
