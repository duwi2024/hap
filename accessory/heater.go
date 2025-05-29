package accessory

import (
	"github.com/duwi2024/hap/service"
)

type Heater struct {
	*A
	Heater *service.Heater
}

// NewHeater returns a heater accessory.
func NewHeater(info Info) *Heater {
	a := Heater{}
	a.A = New(info, TypeHeater)

	a.Heater = service.NewHeater()
	a.AddS(a.Heater.S)

	return &a
}
