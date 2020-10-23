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
	err := models.ReadJSON(request.Body, block)
	if err != nil {
		ResponseError(response, http.StatusBadRequest, err.Error())

		return
	}

	data, err := plc.client.ReadDiscreteInputs(block.Address, block.Quantity)
	if err != nil {
		ResponseError(response, http.StatusInternalServerError, err.Error())

		return
	}

	ResponseOK(response, data, false)
}

func (plc *PLC) readCoils(
	response http.ResponseWriter,
	request *http.Request,
) {
	block := &models.Block{}
	err := models.ReadJSON(request.Body, block)
	if err != nil {
		ResponseError(response, http.StatusBadRequest, err.Error())

		return
	}

	data, err := plc.client.ReadCoils(block.Address, block.Quantity)
	if err != nil {
		ResponseError(response, http.StatusInternalServerError, err.Error())

		return
	}

	ResponseOK(response, data, false)

}

func (plc *PLC) writeCoils(
	response http.ResponseWriter,
	request *http.Request,
) {

}
