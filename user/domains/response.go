package domains

type ResponseJSON struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
