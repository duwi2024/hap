package main

import (
	"context"
	"github.com/LUJUNQUAN/hap"
	"github.com/LUJUNQUAN/hap/accessory"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a := accessory.NewMultiParameterSensor(accessory.Info{
		Name: "MultiParameterSensor",
	})

	s, err := hap.NewServer(hap.NewFsStore("./db"), a.A)
	if err != nil {
		log.Panic(err)
	}

	s.Pin = "34679023"

	a.AirQualityMultiSensor.AirQuality.SetValue(2)
	a.TemperatureSensor.CurrentTemperature.SetValue(34.1)
	a.HumiditySensor.CurrentRelativeHumidity.SetValue(98.3)
	a.AirQualityMultiSensor.PM2_5Density.SetValue(230)
	a.AirQualityMultiSensor.VocDensity.SetValue(120)
	a.AirQualityMultiSensor.PM10Density.SetValue(345)
	a.CarbonDioxideSensor.CarbonDioxideLevel.SetValue(2354)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		signal.Stop(c)
		cancel()
	}()

	s.ListenAndServe(ctx)

}
