package pincode

import (
	"fmt"
	"strconv"
	"strings"
)

// wiringPi Pin  Name  Board Pin  BCM GPIO
// 0  GPIO 0  11  17
// 1  GPIO 1  12  18
// 2  GPIO 2  13  21
// 3  GPIO 3  15  22
// 4  GPIO 4  16  23
// 5  GPIO 5  18  24
// 6  GPIO 6  22  25
// 7  GPIO 7  7  4
// 8  SDA  3  0
// 9  SCL  5  1
// 10  CE0  24  8
// 11  CE1  26  7
// 12  MOSI  19  10
// 13  MISO  21  9
// 14  SCLK  23  11
// 15  TXD  8  14
// 16  RXD  10  15
// 17  GPIO 8    28  //  use 0 space hold
// 18  GPIO 9    29  //
// 19  GPIO10    30  //
// 20  GPIO11    31  //

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

var sheet = `
0  GPIO 0 11  17
1  GPIO 1  12  18
2  GPIO 2  13  21
3  GPIO 3  15  22
4  GPIO 4  16  23
5  GPIO 5  18  24
6  GPIO 6  22  25
7  GPIO 7  7  4
8  SDA  3  0
9  SCL  5  1
10  CE0  24  8
11  CE1  26  7
12  MOSI  19  10
13  MISO  21  9
14  SCLK  23  11
15  TXD  8  14
16  RXD  10  15
17  GPIO 8  0  28
18  GPIO 9  0  29
19  GPIO 10  0  30
20  GPIO 11  0  31
`
var (
	//PinCodeArray
	PinCodeArray = []PinCode{}
	// WiringPiMap
	WiringPiMap = map[uint]PinCode{}
	// BCMMap
	BCMMap = map[uint]PinCode{}
	// BoardMap
	BoardMap = map[uint]PinCode{}
)

func init() {

	lines := strings.Split(sheet, "\n")
	for _, line := range lines {
		tsline := strings.TrimSpace(line)
		if tsline == "" {
			continue
		}
		parts := strings.Fields(tsline)
		var pc PinCode
		var wirecode, bcmcode, boardcode uint
		var name string
		iwcode, _ := strconv.Atoi(parts[0])
		wirecode = uint(iwcode)
		if len(parts) == 4 {
			name = parts[1]
			iboardcode, _ := strconv.Atoi(parts[2])
			boardcode = uint(iboardcode)
			ibmccode, _ := strconv.Atoi(parts[3])
			bcmcode = uint(ibmccode)

		} else if len(parts) == 5 {

			name = fmt.Sprintf("%s %s", parts[1], parts[2])
			iboardcode, _ := strconv.Atoi(parts[3])
			boardcode = uint(iboardcode)
			ibmccode, _ := strconv.Atoi(parts[4])
			bcmcode = uint(ibmccode)
		}
		pc = PinCode{
			WiringPi: wirecode,
			Name:     name,
			BCM:      bcmcode,
			Board:    boardcode,
		}

		PinCodeArray = append(PinCodeArray, pc)
		WiringPiMap[wirecode] = pc
		BCMMap[bcmcode] = pc
		BoardMap[boardcode] = pc
	}
}
