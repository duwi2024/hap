// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/duwi2024/hap/characteristic"
)

type Thermostat_Heater struct {
	*S

	CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
	TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
	CurrentTemperature         *characteristic.CurrentTemperature
	TargetTemperature          *characteristic.TargetTemperature
	TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits
}

func NewThermostat_Heater() *Thermostat_Heater {
	s := Thermostat_Heater{}
	s.S = New(TypeThermostat)

	s.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	s.CurrentHeatingCoolingState.ValidVals = []int{
		characteristic.CurrentHeatingCoolingStateOff,
		characteristic.CurrentHeatingCoolingStateHeat,
	}

	s.AddC(s.CurrentHeatingCoolingState.C)

	s.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	s.TargetHeatingCoolingState.ValidVals = []int{
		characteristic.TargetHeatingCoolingStateOff,
		characteristic.TargetHeatingCoolingStateHeat,
	}
	s.AddC(s.TargetHeatingCoolingState.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	s.TargetTemperature = characteristic.NewTargetTemperature()
	s.AddC(s.TargetTemperature.C)

	s.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	s.AddC(s.TemperatureDisplayUnits.C)

	return &s
}
