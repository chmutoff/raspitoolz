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
		println(i)
		driver.MotorASetSpeed(i)
		driver.MotorBSetSpeed(i)
		time.Sleep(time.Millisecond * 100)
	}
	for i := uint8(100); i >= 15; i-- {
		println(i)
		driver.MotorASetSpeed(i)
		driver.MotorBSetSpeed(i)
		time.Sleep(time.Millisecond * 100)
		if i == 15 {
			driver.MotorASetSpeed(0)
			driver.MotorBSetSpeed(0)
		}
	}

	/*
		driver.MotorASetSpeed(100)
		driver.MotorBSetSpeed(100)

		driver.MotorAClockwise()
		driver.MotorBClockwise()
		time.Sleep(time.Second * 5)

		driver.MotorACounterclockwise()
		driver.MotorBCounterclockwise()
		time.Sleep(time.Second * 5)

		driver.MotorASetSpeed(0)
		driver.MotorBSetSpeed(0)
		time.Sleep(time.Millisecond * 10)
	*/
}
