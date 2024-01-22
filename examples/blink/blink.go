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

	outputs := []uint{rpi.GPIO22, rpi.GPIO23, rpi.GPIO24, rpi.GPIO25} //,  rpi.GPIO25}
	colorLED := []string{"Red", "Yellow", "Green", "Blue"}
	var outPins [4]gpio.Pin
	fmt.Println(outputs)
	for i, outp := range outputs {
		pin, err := gpio.OpenPin(int(outp), gpio.ModeOutput)
		outPins[i] = pin
		if err != nil {
			fmt.Printf("Error opening %d pin = %v, Error: %s\n", outp, pin, err)
			return
		}
	}

	// turn the led off on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nClearing and unexporting the pins.\n")
			for _, pin := range outPins {
				pin.Clear()
				pin.Close()
			}
			os.Exit(0)
		}
	}()

	for j := 1; j <= 8; j++ {
		for k, pin := range outPins {
			pin.Set()
			time.Sleep(700 * time.Millisecond)
			pin.Clear()
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Loop %d Output %s\n", j, colorLED[k])
		}
	}
}
