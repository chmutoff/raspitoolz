package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func main() {
	_ = rpio.Open()
	defer rpio.Close()

	var pin = rpio.Pin(18)
	pin.Mode(rpio.Pwm)
	pin.Freq(1920)
	var t = time.Millisecond * 300

	// Go to all 3 positions and then backwards
	duty := []uint32{5, 15, 25}
	for i := 0; i < 2; i++ {
		for j := 0; j < len(duty); j++ {
			pin.DutyCycle(duty[j], 200)
			time.Sleep(t)
		}
		for j := len(duty) - 1; j >= 0; j-- {
			pin.DutyCycle(duty[j], 200)
			time.Sleep(t)
		}
	}

	// Alternate between first and second position
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			pin.DutyCycle(5, 200)
		} else {
			pin.DutyCycle(15, 200)
		}
		time.Sleep(t)
	}

	// Alternate between third and fourth position
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			pin.DutyCycle(15, 200)
		} else {
			pin.DutyCycle(25, 200)
		}
		time.Sleep(t)
	}

	// Alternate from first to last position
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			pin.DutyCycle(5, 200)
		} else {
			pin.DutyCycle(25, 200)
		}
		time.Sleep(t * 2)
	}

	// Reset to the middle position
	pin.DutyCycle(15, 200)
	time.Sleep(time.Second)

	rpio.StopPwm()
}
