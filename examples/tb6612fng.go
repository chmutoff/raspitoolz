package main

import (
	"github.com/chmutoff/raspitoolz/dcmotordriver"
	"time"
)

func main() {
	driver, _ := dcmotordriver.NewTb6612fng(26, 12, 5, 6, 13, 20, 21)
	defer driver.Close()

	driver.MotorAClockwise()
	driver.MotorBCounterclockwise()
	for i := uint8(15); i <= 100; i++ {
		driver.MotorASetSpeed(i)
		driver.MotorBSetSpeed(i)
		time.Sleep(time.Millisecond * 300)
	}
	time.Sleep(time.Second)
	for i := uint8(100); i >= 15; i-- {
		driver.MotorASetSpeed(i)
		driver.MotorBSetSpeed(i)
		time.Sleep(time.Millisecond * 300)
	}

	driver.MotorASetSpeed(0)
	driver.MotorBSetSpeed(0)
}
