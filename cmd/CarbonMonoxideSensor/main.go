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
	a := accessory.NewCarbonMonoxideSensor(accessory.Info{
		Name: "CarbonMonoxideSensor",
	})

	s, err := hap.NewServer(hap.NewFsStore("./db"), a.A)
	if err != nil {
		log.Panic(err)
	}

	s.Pin = "12569078"

	a.CarbonMonoxideSensor.CarbonMonoxideLevel.SetValue(230)

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
