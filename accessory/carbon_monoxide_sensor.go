package accessory

import "github.com/LUJUNQUAN/hap/service"

type CarbonMonoxideSensor struct {
	*A
	CarbonMonoxideSensor *service.CarbonMonoxideSensor
}

// NewCarbonMonoxideSensor return a CarbonMonoxideSensor which implements model.CarbonMonoxideSensor.
func NewCarbonMonoxideSensor(info Info) *CarbonMonoxideSensor {
	a := CarbonMonoxideSensor{}
	a.A = New(info, TypeSensor)

	a.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	a.AddS(a.CarbonMonoxideSensor.S)

	return &a
}
