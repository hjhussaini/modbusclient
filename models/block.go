package models

type Block struct {
	Address  uint16   `json:"address"`
	Quantity uint16   `json:"quantity"`
	Value    []uint16 `json:"value,omitempty"`
}
