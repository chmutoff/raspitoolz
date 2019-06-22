package main

import (
	"log"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/pca9685"
	"periph.io/x/periph/host"
	"time"
)

// Example of moving two servos from 0 - 180 degrees with PAC9685
// https://github.com/google/periph/tree/master/experimental/devices/pca9685
func Example() {
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	pca, err := pca9685.NewI2C(bus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}

	if err := pca.SetPwmFreq(50 * physic.Hertz); err != nil {
		log.Fatal(err)
	}
	if err := pca.SetAllPwm(0, 0); err != nil {
		log.Fatal(err)
	}
	servos := pca9685.NewServoGroup(pca, 1, 650, 0, 180)

	horizontalServo := servos.GetServo(0)
	verticalServo := servos.GetServo(1)

	for k := 0; k < 3; k++ {
		for i := 0; i <= 180; i++ {
			if err := horizontalServo.SetAngle(physic.Angle(i)); err != nil {
				log.Fatal(err)
			}

			if err := verticalServo.SetAngle(physic.Angle(i)); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond * 3)
		}

		for i := 180; i >= 0; i-- {
			if err := horizontalServo.SetAngle(physic.Angle(i)); err != nil {
				log.Fatal(err)
			}

			if err := verticalServo.SetAngle(physic.Angle(i)); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond * 3)
		}
	}

	if err := horizontalServo.SetAngle(90); err != nil {
		log.Fatal(err)
	}

	if err := verticalServo.SetAngle(90); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Example()
}
