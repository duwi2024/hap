// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/duwi2024/hap/characteristic"
)

const TypeTemperatureSensor = "8A"

type TemperatureSensor struct {
	*S

	CurrentTemperature *characteristic.CurrentTemperature
}

func NewTemperatureSensor() *TemperatureSensor {
	s := TemperatureSensor{}
	s.S = New(TypeTemperatureSensor)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	return &s
}
