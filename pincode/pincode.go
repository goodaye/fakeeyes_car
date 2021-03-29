package pincode

//

// PinCode pin code
type PinCode struct {
	WiringPi uint
	Name     string
	Board    uint
	BCM      uint
}

// Wiring2BCM transfer wiringPi Code to BCM code
func Wiring2BCM(w uint) uint {
	return WiringPiMap[w].BCM
}
