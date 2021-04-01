package clutch

// Led light led
func Led(R uint, G uint, B uint) {

	if R == 0 {
		PinLedR.Low()
	} else {
		PinLedR.High()
	}

	if G == 0 {
		PinLedG.Low()
	} else {
		PinLedG.High()
	}

	if B == 0 {
		PinLedB.Low()
	} else {
		PinLedB.High()
	}
}
