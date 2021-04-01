package clutch

import "fmt"

// MotorLeft turn left
func MotorLeft() {
	fmt.Println("left run ")
	PinLeftmotorgo.High()
	PinLeftmotorback.Low()
	PinLeftmotorpwm.DutyCycle(100, 100)
	// PinLeftmotorpwm.Freq(100000)
	// time.Sleep(100 * time.Microsecond)
	// pinleftmotor.Low()

}
