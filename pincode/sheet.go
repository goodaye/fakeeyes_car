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
