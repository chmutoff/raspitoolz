package shiftregister

import "io"

type BasicShiftRegister interface {
	WriteData(uint)
	GetData() uint
	io.Closer
}
