package servo

import "github.com/stianeikeland/go-rpio/v4"

type mg996r struct {
	sg90
}

func NewMg996R(pin int) (*mg996r, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	servo := mg996r{}
	servo.initPin(pin)

	return &servo, nil
}
