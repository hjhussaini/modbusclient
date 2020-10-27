package api

import "net/http"

func (plc *PLC) readInputRegisters(
	response http.ResponseWriter,
	request *http.Request,
) {

}

func (plc *PLC) readHoldingRegisters(
	response http.ResponseWriter,
	request *http.Request,
) {

}

func (plc *PLC) readRegisters(
	response http.ResponseWriter,
	request *http.Request,
) {

}

func (plc *PLC) writeRegisters(
	response http.ResponseWriter,
	request *http.Request,
) {

}

func (plc *PLC) readFIFOQueue(
	response http.ResponseWriter,
	request *http.Request,
) {

}

func bytesToUint16(array []byte) (results []uint16) {
	for index := 1; index < len(array); index += 2 {
		results = append(results, uint16(array[index])*256+uint16(array[index+1]))
	}

	return results
}
