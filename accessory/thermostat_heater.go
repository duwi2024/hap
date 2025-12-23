package accessory

import (
	"github.com/duwi2024/hap/service"
)

type Thermostat_Heater struct {
	*A
	Thermostat *service.Thermostat_Heater
}

// NewThermostat returns a Thermostat accessory.
func NewThermostat_Heater(info Info) *Thermostat_Heater {
	a := Thermostat_Heater{}
	a.A = New(info, TypeThermostat)

	a.Thermostat = service.NewThermostat_Heater()
	a.AddS(a.Thermostat.S)

	return &a
}
