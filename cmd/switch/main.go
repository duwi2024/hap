// This example show an example of a switch accessory
// which periodically changes it's state between on and off.
package main

import (
	"github.com/LUJUNQUAN/hap"
	"github.com/LUJUNQUAN/hap/accessory"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	a := accessory.NewBridge(accessory.Info{
		Name: "bridge",
	})

	b := accessory.NewSwitch(accessory.Info{
		Name: "Lamp",
	})

	d := accessory.NewSwitch(accessory.Info{
		Name: "Lamp1",
	})

	e := accessory.NewWindowCoveringVertical(accessory.Info{
		Name: "Vertical",
	})

	f := accessory.NewWindowCoveringHorizontal(accessory.Info{
		Name: "Horizontal",
	})

	c1 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-2",
	})

	c2 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-3",
	})

	c3 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-4",
	})

	c4 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-5",
	})

	c5 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-6",
	})

	c6 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-7",
	})

	c7 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-8",
	})

	c8 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-9",
	})

	c9 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-10",
	})

	c10 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-11",
	})

	c11 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-12",
	})

	c12 := accessory.NewSwitch(accessory.Info{
		Name: "Lamp02341232-13",
	})

	h := accessory.NewTelevision(accessory.Info{
		Name: "Television",
	})

	s, err := hap.NewServer(hap.NewFsStore("./db"), a.A, b.A, d.A, e.A, f.A, c1.A, c2.A, c3.A, c4.A, c5.A, c6.A, c7.A, c8.A, c9.A, c10.A, c11.A, c12.A, h.A)
	if err != nil {
		log.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	b.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			log.Println("Client changed switch to on")
		} else {
			log.Println("Client changed switch to off")
		}
	})

	s.Pin = "34679023"

	// Periodically toggle the switch's on characteristic
	go func() {
		for {
			on := !b.Switch.On.Value()
			if on {
				log.Println("Switch is on")
			} else {
				log.Println("Switch is off")
			}
			b.Switch.On.SetValue(on)
			time.Sleep(1 * time.Second)
		}
	}()

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
