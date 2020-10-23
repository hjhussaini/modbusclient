package api

type Result struct {
	Message string   `json:"message"`
	Data    []uint16 `json:"data"`
}
