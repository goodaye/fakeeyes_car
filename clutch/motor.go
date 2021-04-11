package clutch

import "fmt"

// MotorLeft turn left
func MotorRun() {
	fmt.Println("left run ")
	PinLeftMotorgo.High()
	PinLeftMotorback.Low()
	PinLeftMotorpwm.DutyCycle(100, 255)

	PinRightMotorgo.High()
	PinRightMotorback.Low()
	PinRightMotorpwm.DutyCycle(100, 255)
}
