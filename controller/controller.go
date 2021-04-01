package controller

import (
	"time"

	"github.com/mateemu/fakeeyes_car/clutch"
)

// TurnLeft car turn left
func TurnLeft() {

	clutch.TurnLeft()
}

//BeepTimes beep times
func BeepTimes(c int, d time.Duration) {
	for i := 0; i < c; i++ {
		clutch.Beep(d)
		time.Sleep(d)
	}
}

func MotorLeft() {
	clutch.MotorLeft()
}

func MotorRun() {

}

func Led(R uint, G uint, B uint) {
	clutch.Led(R, G, B)
}
