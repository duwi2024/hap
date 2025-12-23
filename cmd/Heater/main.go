// This example show an example of a switch accessory
// which periodically changes it's state between on and off.
package main

import (
	"github.com/duwi2024/hap"
	"github.com/duwi2024/hap/accessory"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	a := accessory.NewBridge(accessory.Info{
		Name: "bridge",
	})

	b := accessory.NewThermostat_Heater(accessory.Info{
		Name: "thermostat",
	})

	s, err := hap.NewServer(hap.NewFsStore("./db"), a.A, b.A)
	if err != nil {
		log.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic

	b.Thermostat.CurrentTemperature.SetMaxValue(60)
	b.Thermostat.CurrentTemperature.SetMinValue(-20)
	b.Thermostat.CurrentTemperature.SetStepValue(0.5)

	b.Thermostat.TargetTemperature.SetMaxValue(45)
	b.Thermostat.TargetTemperature.SetMinValue(0)
	b.Thermostat.TargetTemperature.SetStepValue(1)

	b.Thermostat.TargetTemperature.SetValue(20)
	b.Thermostat.CurrentTemperature.SetValue(18)

	b.Thermostat.CurrentHeatingCoolingState.OnSetRemoteValue(func(v int) error {
		log.Printf("CurrentHeatingCoolingState is %v", v)
		b.Thermostat.CurrentHeatingCoolingState.SetValue(v)
		return nil
	})

	b.Thermostat.TargetHeatingCoolingState.OnSetRemoteValue(func(v int) error {
		log.Printf("TargetHeatingCoolingState is %v", v)
		b.Thermostat.TargetHeatingCoolingState.SetValue(v)
		if v == 1 {
			b.Thermostat.CurrentHeatingCoolingState.SetValue(1)
		} else {
			b.Thermostat.CurrentHeatingCoolingState.SetValue(0)
		}
		return nil
	})

	b.Thermostat.TargetTemperature.OnSetRemoteValue(func(v float64) error {
		log.Printf("TargetTemperature is %v", v)
		b.Thermostat.TargetTemperature.SetValue(v)
		return nil
	})

	b.Thermostat.CurrentTemperature.OnSetRemoteValue(func(v float64) error {
		log.Printf("CurrentTemperature is %v", v)
		b.Thermostat.CurrentTemperature.SetValue(v)
		return nil
	})

	s.Pin = "34679023"

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
