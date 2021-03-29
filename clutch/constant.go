package clutch

import (
	"fmt"

	"github.com/mateemu/fakeeyes_car/pincode"
	"github.com/stianeikeland/go-rpio/v4"
)

/*蜂鸣器引脚设置*/
//设置控制蜂鸣器引脚为wiringPi编码10口
var (
	pinbuzzercode = pincode.Wiring2BCM(10)
	pinbuzzer     rpio.Pin
)

// pin := rpio.Pin(10)

// pin.Output()       // Output mode
// pin.High()         // Set pin High
// pin.Low()          // Set pin Low
// pin.Toggle()       // Toggle pin (Low -> High -> Low)

// pin.Input()        // Input mode
// res := pin.Read()  // Read state from pin (High / Low)

// pin.Mode(rpio.Output)   // Alternative syntax
// pin.Write(rpio.High)    // Alternative syntax

func initpin() {
	pinbuzzer = rpio.Pin(pinbuzzercode)
	fmt.Println(pinbuzzercode)
	pinbuzzer.Output()
}
