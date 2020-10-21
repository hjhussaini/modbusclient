package api

import (
	"net/http"

	"github.com/goburrow/modbus"
	"github.com/gorilla/mux"
)

type PLC struct {
	client modbus.Client
}

func (plc *PLC) RegisterAPIs(router *mux.Router) {
	postRouter := router.Methods(http.MethodPost).Subrouter()

	postRouter.HandleFunc("/connect", plc.connect)
	postRouter.HandleFunc("/disconnect", plc.disconnect)
	postRouter.HandleFunc("/coils", plc.writeCoils)
	postRouter.HandleFunc("/registers", plc.writeRegisters)

	getRouter := router.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/coils/inputs", plc.readDiscreteInputs)
	getRouter.HandleFunc("/coils", plc.readCoils)
	getRouter.HandleFunc("/registers/inputs", plc.readInputRegisters)
	getRouter.HandleFunc("/registers/holdings", plc.readHoldingRegisters)
	getRouter.HandleFunc("/registers", plc.readRegisters)
	getRouter.HandleFunc("/fifo", plc.readFIFOQueue)
}

func NewPLC() *PLC {
	return &PLC{}
}
