package ams

import (
	"github.com/chmutoff/raspitoolz/servo"
	"github.com/chmutoff/raspitoolz/shiftregister"
	"github.com/stianeikeland/go-rpio"
)

type dcMotor struct {
	pin       rpio.Pin
	direction DcMotorDirection
	speed     uint
}

type DcMotorId uint

const (
	DCM1 DcMotorId = iota
	DCM2
	DCM3
	DCM4
)

type stepperMotor struct {
}

type StepperMotorId uint

const (
	STM1 StepperMotorId = iota
	STM2
)

type ServoMotorId uint

const (
	SRM1 ServoMotorId = iota
	SRM2
)

var DcMotorDirectionValue = map[DcMotorId]map[DcMotorDirection]uint{
	DCM1: {
		Clockwise:        4,
		CounterClockwise: 8,
	},
	DCM2: {
		Clockwise:        2,
		CounterClockwise: 16,
	},
	DCM3: {
		Clockwise:        1,
		CounterClockwise: 64,
	},
	DCM4: {
		Clockwise:        32,
		CounterClockwise: 128,
	},
}

type shield struct {
	ctrl     shiftregister.BasicShiftRegister
	dcMotors [4]dcMotor
	stMotors [2]stepperMotor
	srMotors [2]servo.Sg90
}

type DcMotorDirection uint8

const (
	None DcMotorDirection = iota
	Clockwise
	CounterClockwise
)

func NewShield(ser int, srclk int, rclk int) (*shield, error) {
	sreg, _ := shiftregister.NewSn74hc595(ser, srclk, rclk, 8, true)
	shield := shield{ctrl: sreg}
	_ = rpio.Open()
	return &shield, nil
}

func (shield *shield) AddDcMotor(motor DcMotorId, pin int) {
	shield.dcMotors[motor].pin = rpio.Pin(pin)

	// TODO: check if pin is pwm and if it's an individual pin
	shield.dcMotors[motor].pin.Pwm()
	shield.dcMotors[motor].pin.Freq(1920)
}

/*
func (shield *shield) StartDcMotor(motor DcMotorId, direction DcMotorDirection, speed uint) {
}

func (shield *shield) StopDcMotor(motor DcMotorId) {
}
*/

func (shield *shield) SetDcMotorSpeedAndDirection(id DcMotorId, speed uint, direction DcMotorDirection) {
	shield.SetDcMotorSpeed(id, speed)
	shield.SetDcMotorDirection(id, direction)
}

func (shield *shield) SetDcMotorSpeed(id DcMotorId, speed uint) {
	shield.dcMotors[id].pin.DutyCycle(uint32(speed), 100)
	shield.dcMotors[id].speed = speed // TODO: The speed of the pin in the same channel will also change
}

func (shield *shield) SetDcMotorDirection(id DcMotorId, direction DcMotorDirection) {
	currentDirection := shield.dcMotors[id].direction
	if currentDirection != direction { // TODO: maybe a soft change? stop motor and start?
		currentControlValue := DcMotorDirectionValue[id][currentDirection]
		newControlValue := shield.ctrl.GetData() - currentControlValue + DcMotorDirectionValue[id][direction]
		shield.ctrl.WriteData(newControlValue)
		shield.dcMotors[id].direction = direction
	}
}

func (shield shield) AddStepperMotor(id StepperMotorId) {
}

func (shield shield) AddServo(id ServoMotorId, pin int) {
}

func (shield *shield) StopAllMotors() {
	shield.ctrl.WriteData(0)
	shield.dcMotors[DCM1].direction = None
	shield.dcMotors[DCM2].direction = None
	shield.dcMotors[DCM3].direction = None
	shield.dcMotors[DCM4].direction = None
}

func (shield shield) Close() {
	_ = shield.ctrl.Close()
	_ = rpio.Close()
}
