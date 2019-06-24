package dcmotordriver

import "github.com/stianeikeland/go-rpio/v4"

type tb6612fng struct {
	stby rpio.Pin

	pwmA rpio.Pin
	inA1 rpio.Pin
	inA2 rpio.Pin
	minA uint8
	maxA uint8

	pwmB rpio.Pin
	inB1 rpio.Pin
	inB2 rpio.Pin
	minB uint8
	maxB uint8
}

type MotorDirection uint8

const (
	None MotorDirection = iota
	Clockwise
	Counterclockwise
)

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
	driver.pwmA.Freq(1920)

	driver.inA1 = rpio.Pin(aIn1)
	driver.inA1.Output()
	driver.inA1.Low()

	driver.inA2 = rpio.Pin(aIn2)
	driver.inA2.Output()
	driver.inA2.Low()

	driver.pwmB = rpio.Pin(pwmB)
	driver.pwmB.Pwm()
	driver.pwmB.Freq(1920)

	driver.inB1 = rpio.Pin(bIn1)
	driver.inB1.Output()
	driver.inB1.Low()

	driver.inB2 = rpio.Pin(bIn2)
	driver.inB2.Output()
	driver.inB2.Low()

	driver.stby.High() // Enable board

	return &driver, nil
}

func (driver tb6612fng) MotorASetSpeed(speed uint8) {
	driver.pwmA.DutyCycle(uint32(speed), 100)
}

func (driver tb6612fng) MotorAClockwise() {
	driver.MotorASetDirection(Clockwise)
}

func (driver tb6612fng) MotorACounterclockwise() {
	driver.MotorASetDirection(Counterclockwise)
}

func (driver tb6612fng) MotorASetDirection(direction MotorDirection) {
	if direction == Clockwise {
		driver.inA1.High()
		driver.inA2.Low()
	} else {
		driver.inA1.Low()
		driver.inA2.High()
	}
}

func (driver tb6612fng) MotorBSetSpeed(speed uint8) {
	driver.pwmB.DutyCycle(uint32(speed), 100)
}

func (driver tb6612fng) MotorBClockwise() {
	driver.MotorBSetDirection(Clockwise)
}

func (driver tb6612fng) MotorBCounterclockwise() {
	driver.MotorBSetDirection(Counterclockwise)
}

func (driver tb6612fng) MotorBSetDirection(direction MotorDirection) {
	if direction == Clockwise {
		driver.inB1.High()
		driver.inB2.Low()
	} else {
		driver.inB1.Low()
		driver.inB2.High()
	}
}

func (driver tb6612fng) SetMinMaxMotorASpeed(min, max uint8) {

}

func (driver tb6612fng) SetMinMaxMotorBSpeed(min, max uint8) {

}

func (driver tb6612fng) Close() error {
	driver.stby.Low() // Shut down board
	return rpio.Close()
}
