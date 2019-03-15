package dcmotordriver

import "github.com/stianeikeland/go-rpio/v4"

type tb6612fng struct {
	stby rpio.Pin

	pwmA rpio.Pin
	aIn1 rpio.Pin
	aIn2 rpio.Pin

	pwmB rpio.Pin
	bIn1 rpio.Pin
	bIn2 rpio.Pin
}

func NewTb6612fng(stby rpio.Pin, pwmA rpio.Pin, aIn1 rpio.Pin, aIn2 rpio.Pin, pwmB rpio.Pin, bIn1 rpio.Pin, bIn2 rpio.Pin) (*tb6612fng, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	driver := tb6612fng{}

	driver.stby = rpio.Pin(stby)
	driver.stby.Output()
	driver.stby.Low()

	driver.pwmA = rpio.Pin(pwmA)
	driver.pwmA.Pwm()

	driver.aIn1 = rpio.Pin(aIn1)
	driver.aIn1.Output()
	driver.aIn1.Low()

	driver.aIn2 = rpio.Pin(aIn2)
	driver.aIn2.Output()
	driver.aIn2.Low()

	driver.pwmB = rpio.Pin(pwmB)
	driver.pwmB.Pwm()

	driver.bIn1 = rpio.Pin(bIn1)
	driver.bIn1.Output()
	driver.bIn1.Low()

	driver.bIn2 = rpio.Pin(bIn2)
	driver.bIn2.Output()
	driver.bIn2.Low()

	driver.stby.High() // Enable board

	return &driver, nil
}

func (driver tb6612fng) Close() error {
	driver.stby.Low() // Shut down board
	return rpio.Close()
}
