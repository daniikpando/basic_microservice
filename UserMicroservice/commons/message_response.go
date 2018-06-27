package commons

type Message struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Content interface{} `json:"content"`
}