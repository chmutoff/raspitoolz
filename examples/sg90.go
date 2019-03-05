package main

import (
	"github.com/chmutoff/raspitoolz/servo"
	"time"
)

func main() {

	sg90, _ := servo.NewSg90(18)
	defer sg90.Close()

	sg90.InitialPosition()
	time.Sleep(time.Second)

	sg90.NeutralPosition()
	time.Sleep(time.Second)

	sg90.FinalPosition()
	time.Sleep(time.Second)

	sg90.InitialPosition()
	time.Sleep(time.Second)

	for k := 0; k < 5; k++ {
		for i := 0; i <= 180; i++ {
			sg90.Turn(i)
			time.Sleep(time.Millisecond * 5)
			if i == 180 {
				time.Sleep(time.Millisecond * 100)
			}
		}
		for i := 180; i >= 0; i-- {
			sg90.Turn(i)
			time.Sleep(time.Millisecond * 5)
			if i == 180 {
				time.Sleep(time.Millisecond * 100)
			}
		}
	}

	sg90.NeutralPosition()
	time.Sleep(time.Second)
}
