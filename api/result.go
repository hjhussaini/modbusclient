package api

type Result struct {
	Data    []byte `json:"data"`
	Message string `json:"message"`
}
