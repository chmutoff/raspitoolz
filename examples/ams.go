package main

import (
	"github.com/chmutoff/raspitoolz/ams"
	"time"
)

func main() {
	shield, _ := ams.NewShield(17, 27, 22)
	defer shield.Close()

	/* Example 1 */
	shield.AddDcMotor(ams.DCM1, 18)
	shield.AddDcMotor(ams.DCM2, 13)

	for i := uint(20); i <= 100; i += 10 {
		shield.SetDcMotorSpeedAndDirection(ams.DCM1, i, ams.Clockwise)
		time.Sleep(time.Millisecond * 10)
		shield.SetDcMotorSpeedAndDirection(ams.DCM2, i, ams.CounterClockwise)
		time.Sleep(time.Second)
	}

	for i := uint(100); i >= 20; i -= 10 {
		shield.SetDcMotorSpeedAndDirection(ams.DCM1, i, ams.Clockwise)
		time.Sleep(time.Millisecond * 10)
		shield.SetDcMotorSpeedAndDirection(ams.DCM2, i, ams.CounterClockwise)
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 2)
	shield.StopAllMotors()
}
