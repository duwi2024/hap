package accessory

import "github.com/LUJUNQUAN/hap/service"

type MultiParameterSensor struct {
	*A
	TemperatureSensor     *service.TemperatureSensor
	HumiditySensor        *service.HumiditySensor
	AirQualityMultiSensor *service.AirQualityMultiSensor
	CarbonDioxideSensor   *service.CarbonDioxideSensor
}

func NewMultiParameterSensor(info Info) *MultiParameterSensor {
	a := MultiParameterSensor{}
	a.A = New(info, TypeSensor)

	a.TemperatureSensor = service.NewTemperatureSensor()
	a.AddS(a.TemperatureSensor.S)

	a.HumiditySensor = service.NewHumiditySensor()
	a.AddS(a.HumiditySensor.S)

	a.AirQualityMultiSensor = service.NewAirQualityMultiSensor()
	a.AddS(a.AirQualityMultiSensor.S)

	a.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	a.AddS(a.CarbonDioxideSensor.S)

	return &a
}
