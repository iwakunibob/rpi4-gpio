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

	pin22, err := gpio.OpenPin(rpi.GPIO22, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}
	pin23, err := gpio.OpenPin(rpi.GPIO23, gpio.ModeOutput)
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
			pin22.Clear()
			pin23.Clear()
			pin22.Close()
			pin23.Close()
			os.Exit(0)
		}
	}()

	for j := 1; j <= 15; j++ {
		pin22.Set()
		pin23.Clear()
		time.Sleep(200 * time.Millisecond)
		pin23.Set()
		pin22.Clear()
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Blink %v\n", j)
	}
}
