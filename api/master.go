package api

import (
	"net/http"

	"github.com/goburrow/modbus"
	"github.com/hjhussaini/plc/models"
)

func (plc *PLC) connect(
	response http.ResponseWriter,
	request *http.Request,
) {
	device := &models.Device{}
	err := models.ReadJSON(request.Body, device)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		models.WriteJSON(Result{Message: err.Error()}, response)

		return
	}

	switch device.Protocol {
	case "tcp":
		plc.client = modbus.TCPClient(device.Address)
	case "rtu":
		plc.client = modbus.RTUClient(device.Address)
	}

}

func (plc *PLC) disconnect(
	response http.ResponseWriter,
	request *http.Request,
) {

}
