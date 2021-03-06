package shiftregister

import (
	"errors"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
)

// sn74hc595 is a structure for managing an SN74HC595 shift register
type sn74hc595 struct {
	ser           rpio.Pin
	srclk         rpio.Pin
	rclk          rpio.Pin
	outputPins    uint
	positiveLogic bool
	lastData      uint
}

// SetInputPins sets and initializes pins connected to a shift register
func NewSn74hc595(ser int, srclk int, rclk int, outputPins uint, positiveLogic bool) (*sn74hc595, error) {
	sr := sn74hc595{}
	if err := sr.initBoardPins(ser, srclk, rclk); err != nil {
		return nil, err
	}

	if err := sr.initChipPins(outputPins, positiveLogic); err != nil {
		return nil, err
	}

	return &sr, nil
}

func (sr *sn74hc595) initBoardPins(ser int, srclk int, rclk int) error {
	if err := rpio.Open(); err != nil {
		return err
	}

	sr.ser = rpio.Pin(ser)
	sr.ser.Output()
	sr.ser.Low()

	sr.srclk = rpio.Pin(srclk)
	sr.srclk.Output()
	sr.srclk.Low()

	sr.rclk = rpio.Pin(rclk)
	sr.rclk.Output()
	sr.rclk.Low()

	return nil
}

// SetOutputParams defines what kind of characteristics has a connected shift register
func (sr *sn74hc595) initChipPins(outputPins uint, positiveLogic bool) (err error) {
	if outputPins < 8 || outputPins%8 != 0 {
		return errors.New("wrong number of output pins")
	}
	sr.outputPins = outputPins
	sr.positiveLogic = positiveLogic
	return nil
}

// WriteBit writes one bit into shift register
func (sr sn74hc595) writeBit(bit uint) { // ?? bool
	if bit == 0 {
		if sr.positiveLogic == true {
			sr.ser.Low()
		} else {
			sr.ser.High()
		}
	} else {
		if sr.positiveLogic == true {
			sr.ser.High()
		} else {
			sr.ser.Low()
		}
	}
	sr.srclk.High()
	sr.srclk.Low()
}

// Latch moves the data from internal register to memory
func (sr sn74hc595) latch() {
	sr.rclk.High()
	sr.rclk.Low()
}

func (sr *sn74hc595) WriteData(data uint) {
	sr.lastData = data
	var mask = uint(math.Pow(2, float64(sr.outputPins-1)))
	for i := uint(0); i < sr.outputPins; i++ {
		sr.writeBit(mask & (data << i))
	}
	sr.latch()
}

// Returns the last written data
func (sr sn74hc595) GetData() uint {
	return sr.lastData
}

func (sn74hc595) Close() error {
	return rpio.Close()
}
