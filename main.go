package main

import (
        "machine"
        "time"
)

type Led struct {
	machine.Pin
}

func InitLed(pinPosition int) Led {
	led := Led{machine.Pin(pinPosition)}
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return led
}

func (l *Led) Toggle() {
	if l.Get() {
		l.Low()
	} else {
		l.High()
	}
}

func main() {
	ledRed := InitLed(2)
	ledGreen := InitLed(3)
	ledYellow := InitLed(4)

	ledRed.Toggle()
	ledGreen.Toggle()
	ledYellow.Toggle()
	
	time.Sleep(time.Millisecond * 5000)

	for {
		time.Sleep(time.Millisecond * 500)
		ledRed.Toggle()

		time.Sleep(time.Millisecond * 500)
		ledRed.Toggle()
		ledGreen.Toggle()

		time.Sleep(time.Millisecond * 500)
		ledRed.Toggle()

		time.Sleep(time.Millisecond * 500)
		ledRed.Toggle()
		ledGreen.Toggle()
		ledYellow.Toggle()
	}
}
