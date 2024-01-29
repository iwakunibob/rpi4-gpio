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
	// Configure Digital Outputs
	outputs := []uint{rpi.GPIO16, rpi.GPIO17, rpi.GPIO22,
		rpi.GPIO23, rpi.GPIO24, rpi.GPIO25, rpi.GPIO26, rpi.GPIO27}
	colorLED := []string{"Red", "Green", "Blue", "Yellow", "Red", "Green", "Blue", "Yellow"}
	var outputPins [8]gpio.Pin
	fmt.Println(outputs)
	for i, outp := range outputs {
		pin, err := gpio.OpenPin(int(outp), gpio.ModeOutput)
		outputPins[i] = pin
		if err != nil {
			fmt.Printf("Error opening %d pin = %v, Error: %s\n", outp, pin, err)
			return
		}
	}
	// Configure Digital Inputs
	inputs := []uint{rpi.GPIO05, rpi.GPIO06}
	var inputPins [2]gpio.Pin
	fmt.Println(inputs)
	for i, inp := range inputs {
		pin, err := gpio.OpenPin(int(inp), gpio.ModeInput)
		inputPins[i] = pin
		if err != nil {
			fmt.Printf("Error opening %d pin = %v, Error: %s\n", inp, pin, err)
			return
		}
	}

	// Turn the outputs off on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			for _, pin := range outputPins {
				pin.Clear()
				pin.Close()
			}
			for _, pin := range inputPins {
				pin.Close()
			}
			fmt.Printf("\nClearing and reset pins\nSafe Exit\n")
			time.Sleep(500 * time.Millisecond)
			os.Exit(0)
		}
	}()

	// Read Inputs
	for m, pin := range inputPins {
		err := pin.BeginWatch(gpio.EdgeFalling, func() {
			fmt.Println("Input GPIO", pin, "triggered!")
		})
		if err != nil {
			fmt.Printf("Unable to watch input %v: %s\n", inputs[m], err.Error())
			os.Exit(1)
		}
		fmt.Printf("Now watching %v on a falling edge.\n", inputs[m])
	}

	// Looping Output Sequence
	for j := 1; j <= 10; j++ {
		for k, pin := range outputPins {
			fmt.Printf("Loop %d Output %s\n", j, colorLED[k])
			pin.Set()
			time.Sleep(950 * time.Millisecond)
			pin.Clear()
			time.Sleep(50 * time.Millisecond)
		}
	}
}
