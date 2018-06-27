package commons


// Message es la estructura de la respuesta que se quiere enviar al cliente
type Message struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Content interface{} `json:"content"`
}