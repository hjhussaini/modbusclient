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
		ResponseError(response, http.StatusBadRequest, err.Error())

		return
	}

	switch device.Protocol {
	case "tcp":
		plc.handler = modbus.NewTCPClientHandler(device.Address)
	case "rtu":
		plc.handler = modbus.NewRTUClientHandler(device.Address)
	default:
		ResponseError(response, http.StatusBadRequest, "Unknown protocol")

		return
	}
	if err != nil {
		ResponseError(response, http.StatusInternalServerError, err.Error())

		return
	}

	plc.client = modbus.NewClient(plc.handler)

}

func (plc *PLC) disconnect(
	response http.ResponseWriter,
	request *http.Request,
) {
}
