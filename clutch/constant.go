package clutch

import (
	"github.com/stianeikeland/go-rpio/v4"
)

// > gpio readall
// +-----+-----+---------+------+---+---Pi 4B--+---+------+---------+-----+-----+
// | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
// +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
// |     |     |    3.3v |      |   |  1 || 2  |   |      | 5v      |     |     |
// |   2 |   8 |   SDA.1 |  OUT | 1 |  3 || 4  |   |      | 5v      |     |     |
// |   3 |   9 |   SCL.1 |   IN | 1 |  5 || 6  |   |      | 0v      |     |     |
// |   4 |   7 | GPIO. 7 |   IN | 1 |  7 || 8  | 1 | ALT0 | TxD     | 15  | 14  |
// |     |     |      0v |      |   |  9 || 10 | 1 | ALT0 | RxD     | 16  | 15  |
// |  17 |   0 | GPIO. 0 |   IN | 1 | 11 || 12 | 1 | IN   | GPIO. 1 | 1   | 18  |
// |  27 |   2 | GPIO. 2 |  OUT | 0 | 13 || 14 |   |      | 0v      |     |     |
// |  22 |   3 | GPIO. 3 |  OUT | 0 | 15 || 16 | 0 | OUT  | GPIO. 4 | 4   | 23  |
// |     |     |    3.3v |      |   | 17 || 18 | 0 | OUT  | GPIO. 5 | 5   | 24  |
// |  10 |  12 |    MOSI | ALT0 | 0 | 19 || 20 |   |      | 0v      |     |     |
// |   9 |  13 |    MISO |  OUT | 0 | 21 || 22 | 0 | IN   | GPIO. 6 | 6   | 25  |
// |  11 |  14 |    SCLK |  OUT | 0 | 23 || 24 | 1 | OUT  | CE0     | 10  | 8   |
// |     |     |      0v |      |   | 25 || 26 | 0 | IN   | CE1     | 11  | 7   |
// |   0 |  30 |   SDA.0 |  OUT | 0 | 27 || 28 | 0 | OUT  | SCL.0   | 31  | 1   |
// |   5 |  21 | GPIO.21 |   IN | 1 | 29 || 30 |   |      | 0v      |     |     |
// |   6 |  22 | GPIO.22 |   IN | 0 | 31 || 32 | 1 | IN   | GPIO.26 | 26  | 12  |
// |  13 |  23 | GPIO.23 |  OUT | 0 | 33 || 34 |   |      | 0v      |     |     |
// |  19 |  24 | GPIO.24 |  OUT | 0 | 35 || 36 | 0 | OUT  | GPIO.27 | 27  | 16  |
// |  26 |  25 | GPIO.25 |  OUT | 0 | 37 || 38 | 0 | OUT  | GPIO.28 | 28  | 20  |
// |     |     |      0v |      |   | 39 || 40 | 0 | OUT  | GPIO.29 | 29  | 21  |
// +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
// | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
// +-----+-----+---------+------+---+---Pi 4B--+---+------+---------+-----+-----+

var (
	//定义引脚
	/*蜂鸣器引脚设置*/
	//设置控制蜂鸣器引脚为wiringPi编码10口
	PinCodeBuzzer        = 8
	PinCodeLeftmotorgo   = 20
	PinCodeLeftmotorback = 21
	PinCodeLeftmotorpwm  = 16
	PinCodeLedR          = 22
	PinCodeLedG          = 27
	PinCodeLedB          = 24
	//左电机前进AIN2连接Raspberry的wiringPi编码28口
	// Left_motor_go int = 28
	//左电机后退AIN1连接Raspberry的wiringPi编码29口
	// Left_motor_back int = 29
	// // //左电机控速PWMA连接Raspberry的wiringPi编码27口
	// Left_motor_pwm int = 27

	// //右电机前进BIN2连接Raspberry的wiringPi编码24口
	// Right_motor_go int = 24
	// //右电机后退BIN1连接Raspberry的wiringPi编码25口
	// Right_motor_back int = 25

	// //右电机控速PWMB连接Raspberry的wiringPi编码23口
	// Right_motor_pwm int = 23
)
var (
	// PinBuzzer beep
	PinBuzzer        rpio.Pin
	PinLeftmotorgo   rpio.Pin
	PinLeftmotorback rpio.Pin
	PinLeftmotorpwm  rpio.Pin
	PinLedR          rpio.Pin
	PinLedG          rpio.Pin
	PinLedB          rpio.Pin
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
	// Buzzer
	PinBuzzer = rpio.Pin(PinCodeBuzzer)
	PinBuzzer.Output()
	// Motor
	PinLeftmotorgo = rpio.Pin(PinCodeLeftmotorgo)
	PinLeftmotorgo.Output()
	PinLeftmotorback = rpio.Pin(PinCodeLeftmotorback)
	PinLeftmotorback.Output()
	PinLeftmotorpwm = rpio.Pin(PinCodeLeftmotorpwm)
	PinLeftmotorpwm.Pwm()
	PinLeftmotorpwm.Freq(64000)
	// Led
	PinLedR = rpio.Pin(PinCodeLedR)
	PinLedR.Output()
	PinLedG = rpio.Pin(PinCodeLedG)
	PinLedG.Output()
	PinLedB = rpio.Pin(PinCodeLedB)
	PinLedB.Output()
}
