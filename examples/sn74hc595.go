package main

import (
	"github.com/chmutoff/raspitoolz/shiftregister"
	"log"
	"time"
)

func main() {
	sreg, err := shiftregister.NewSn74hc595(17, 27, 18, 16, false)
	if err != nil {
		log.Fatalln(err)
	}

	defer sreg.Close()

	leds := []uint{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x0100, 0x0200, 0x0400, 0x0800, 0x1000, 0x2000, 0x4000, 0x8000}
	for {
		for _, led := range leds {
			sreg.Write(led)
			sreg.Latch()
			time.Sleep(time.Millisecond * 70)
		}
	}
}
