package api

import (
	"net/http"

	"github.com/hjhussaini/plc/models"
)

func ResponseError(response http.ResponseWriter, status int, message string) {
	response.WriteHeader(status)
	models.WriteJSON(Result{Message: message}, response)
}

func ResponseOK(
	response http.ResponseWriter,
	results []byte,
	is16Bits bool,
) {
	data := make([]uint16, len(results))
	for index, _ := range results {
		if is16Bits {
		} else {
			data[index] = uint16(results[index])
		}
	}
	models.WriteJSON(Result{Message: "OK", Data: data}, response)
}
