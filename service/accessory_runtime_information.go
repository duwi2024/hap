package service

import "github.com/duwi2024/hap/characteristic"

const TypeAccessoryRuntimeInformation = "239"

type AccessoryRuntimeInformation struct {
	*S

	Ping *characteristic.Ping
}

func NewAccessoryRuntimeInformation() *AccessoryRuntimeInformation {
	s := AccessoryRuntimeInformation{}
	s.S = New(TypeAccessoryRuntimeInformation)

	s.Ping = characteristic.NewPing()
	s.AddC(s.Ping.C)

	return &s
}
