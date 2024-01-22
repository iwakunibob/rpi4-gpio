package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/iwakunibob/gpio"
	"github.com/iwakunibob/gpio/rpi"
)

func main() {
	// set GPIO25 to output mode
	pin, err := gpio.OpenPin(rpi.GPIO22, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	// turn the led off on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nClearing and unexporting the pin.\n")
			pin.Clear()
			pin.Close()
			os.Exit(0)
		}
	}()

	for i := 1; i <= 16; i++ {
		pin.Set()
		time.Sleep(200 * time.Millisecond)
		pin.Clear()
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Blink %v\n", i)
	}
}
