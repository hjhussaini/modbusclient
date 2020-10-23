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
		plc.handler = modbus.NewTCPClientHandler(device.Address)
	case "rtu":
		plc.handler = modbus.NewRTUClientHandler(device.Address)
	default:
		return
	}
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		models.WriteJSON(Result{Message: err.Error()}, response)

		return
	}

	plc.client = modbus.NewClient(plc.handler)

}

func (plc *PLC) disconnect(
	response http.ResponseWriter,
	request *http.Request,
) {
}
