package api

import (
	"net/http"

	"github.com/hjhussaini/plc/models"
)

func (plc *PLC) readDiscreteInputs(
	response http.ResponseWriter,
	request *http.Request,
) {
	block := &models.Block{}
	err := models.ReadJSON(request.Body, &block)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		models.WriteJSON(Result{Message: err.Error()}, response)

		return
	}

	data, err := plc.client.ReadDiscreteInputs(block.Address, block.Quantity)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		models.WriteJSON(Result{Message: err.Error()}, response)

		return
	}

	models.WriteJSON(Result{Data: data}, response)
}

func (plc *PLC) readCoils(
	response http.ResponseWriter,
	request *http.Request,
) {
	block := &models.Block{}
	err := models.ReadJSON(request.Body, block)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		models.WriteJSON(Result{Message: err.Error()}, response)

		return
	}

	data, err := plc.client.ReadCoils(block.Address, block.Quantity)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		models.WriteJSON(Result{Message: err.Error()}, response)

		return
	}

	models.WriteJSON(Result{Data: data}, response)

}

func (plc *PLC) writeCoils(
	response http.ResponseWriter,
	request *http.Request,
) {

}
