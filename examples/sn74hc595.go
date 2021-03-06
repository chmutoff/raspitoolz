package main

import (
	"github.com/chmutoff/raspitoolz/shiftregister"
	"log"
	"time"
)

func main() {
	sreg, err := shiftregister.NewSn74hc595(17, 27, 18, 8, false)
	if err != nil {
		log.Fatalln(err)
	}

	defer sreg.Close()

	leds := []uint{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}
	for {
		for _, led := range leds {
			sreg.Write(led)
			sreg.Latch()
			time.Sleep(time.Millisecond * 70)
		}
	}
}
