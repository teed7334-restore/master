package bots

type curl interface {
	Post(url string, params []byte) []byte
	PostForm(url string, params string, header map[string]string) []byte
}
