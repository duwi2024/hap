package service

import "github.com/LUJUNQUAN/hap/characteristic"

const TypeAirQualityMultiSensor = "8D"

type AirQualityMultiSensor struct {
	*S

	AirQuality   *characteristic.AirQuality
	VocDensity   *characteristic.VOCDensity
	PM2_5Density *characteristic.PM2_5Density
	PM10Density  *characteristic.PM10Density
}

func NewAirQualityMultiSensor() *AirQualityMultiSensor {
	s := AirQualityMultiSensor{}
	s.S = New(TypeAirQualityMultiSensor)

	s.AirQuality = characteristic.NewAirQuality()
	s.AddC(s.AirQuality.C)

	s.VocDensity = characteristic.NewVOCDensity()
	s.AddC(s.VocDensity.C)

	s.PM2_5Density = characteristic.NewPM2_5Density()
	s.AddC(s.PM2_5Density.C)

	s.PM10Density = characteristic.NewPM10Density()
	s.AddC(s.PM10Density.C)

	return &s
}
