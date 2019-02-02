package servo

import (
	"github.com/stianeikeland/go-rpio/v4"
)

const cycleLength = 400

type Sg90 struct {
	pin rpio.Pin
}

func NewSg90(pin int) (*Sg90, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	servo := Sg90{}
	servo.initPin(pin)

	return &servo, nil
}

func (servo *Sg90) initPin(pin int) {
	servo.pin = rpio.Pin(pin)
	servo.pin.Pwm()
	servo.pin.Freq(20000) //50 * cycleLength
}

func (servo Sg90) InitialPosition() {
	servo.pin.DutyCycle(10, cycleLength)
}

func (servo Sg90) NeutralPosition() {
	servo.pin.DutyCycle(30, cycleLength)
}

func (servo Sg90) FinalPosition() {
	servo.pin.DutyCycle(50, cycleLength)
}

func (servo Sg90) Turn(degree int) {
	y := (4.0 / 18.0 * float64(degree)) + 10.0
	servo.pin.DutyCycle(uint32(y), cycleLength)
}

func (Sg90) Close() error {
	return rpio.Close()
}
