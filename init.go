package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Event matching struct
type evMatch struct {
	Code, Type uint16
	Value      int32
}

//Initialize match events
// Right Joycon events
var btnXEvent = evMatch{
	Code: 305, Type: 1,
}
var btnYEvent = evMatch{
	Code: 307, Type: 1,
}
var btnAEvent = evMatch{
	Code: 304, Type: 1,
}
var btnBEvent = evMatch{
	Code: 306, Type: 1,
}
var btnPlusEvent = evMatch{
	Code: 313, Type: 1,
}
var btnHomeEvent = evMatch{
	Code: 316, Type: 1,
}
var btnREvent = evMatch{
	Code: 318, Type: 1,
}
var btnRZEvent = evMatch{
	Code: 319, Type: 1,
}
var btnSLEvent = evMatch{
	Code: 308, Type: 1,
}
var btnSREvent = evMatch{
	Code: 309, Type: 1,
}
var rightStickHEvent = evMatch{
	Code: 17, Type: 3,
}
var rightStickVEvent = evMatch{
	Code: 16, Type: 3,
}
var btnR3Event = evMatch{
	Code: 315, Type: 1,
}
// Left Joycon events
var btnDownEvent = evMatch{
	Code: 305, Type: 1,
}
var btnRightEvent = evMatch{
	Code: 307, Type: 1,
}
var btnLeftEvent = evMatch{
	Code: 304, Type: 1,
}
var btnUpEvent = evMatch{
	Code: 306, Type: 1,
}
var btnMinusEvent = evMatch{
	Code: 312, Type: 1,
}
var btnSSEvent = evMatch{
	Code: 317, Type: 1,
}
var btnLEvent = evMatch{
	Code: 318, Type: 1,
}
var btnLZEvent = evMatch{
	Code: 319, Type: 1,
}
var leftStickHEvent = evMatch{
	Code: 17, Type: 3,
}
var leftStickVEvent = evMatch{
	Code: 16, Type: 3,
}
var btnL3Event = evMatch{
	Code: 314, Type: 1,
}
// Other events
var endFrameEvent = evMatch{
	Code: 0, Type: 0,
}

// Global button vars
var btnA, btnB string
var btnX, btnY string
var btnPlus, btnMinus string
var btnUp, btnDown string
var btnLeft, btnRight string
var btnL, btnLZ, btnL3 string
var btnR, btnRZ, btnR3 string
var btnHome, btnScreenShot string
var btnRightSL, btnRightSR string
var btnLeftSL, btnLeftSR string
var lsLeft, lsRight, lsUp, lsDown string
var rsLeft, rsRight, rsUp, rsDown string

func init () {
	fmt.Println("-------------JoyGo Start-------------")
	if len(os.Args) >= 2 {
		// If file exists
		if _, err := os.Stat(os.Args[1]); err == nil {
			loadKeyMap()
		} else {
			fmt.Println("Specified config file does not exist!")
			os.Exit(0)
		}
	} else {
		fmt.Println("No config file specified!")
		os.Exit(0)
	}
}

func loadKeyMap () {
	// Open config file for reading
	config, err := os.Open(os.Args[1])
	parseFatal(err, "Error opening config file!")
	defer config.Close()

	// Make scanner for line by line read
	scanner := bufio.NewScanner(config)
	// Read through file and set joycon/button config
	for scanner.Scan() {
		btnVals := strings.Split(scanner.Text(), "=")
		switch btnVals[0] {
			case "joycons":
				if btnVals[1] == "rl" || btnVals[1] == "lr" {
					joyconR = true
					joyconL = true
				} else if btnVals[1] == "r" {
					joyconR = true
					joyconL = false
				} else if btnVals[1] == "l" {
					joyconR = false
					joyconL = true
				} else if btnVals[1] == "" {
					fmt.Println("You must specify at least one joycon!")
					os.Exit(0)
				} else {
					fmt.Println("Incorrect joycon specification!")
					fmt.Println("Use r, l, rl, or lr.")
					os.Exit(0)
				}
			case "btnA":
				btnA = btnVals[1]
			case "btnB":
				btnB = btnVals[1]
			case "btnX":
				btnX = btnVals[1]
			case "btnY":
				btnY = btnVals[1]
			case "btnPlus":
				btnPlus = btnVals[1]
			case "btnMinus":
				btnMinus = btnVals[1]
			case "btnUp":
				btnUp = btnVals[1]
			case "btnDown":
				btnDown = btnVals[1]
			case "btnLeft":
				btnLeft = btnVals[1]
			case "btnRight":
				btnRight = btnVals[1]
			case "btnL":
				btnL = btnVals[1]
			case "btnLZ":
				btnLZ = btnVals[1]
			case "btnL3":
				btnL3 = btnVals[1]
			case "btnR":
				btnR = btnVals[1]
			case "btnRZ":
				btnRZ = btnVals[1]
			case "btnR3":
				btnR3 = btnVals[1]
			case "btnHome":
				btnHome = btnVals[1]
			case "btnScreenShot":
				btnScreenShot = btnVals[1]
			case "btnRightSL":
				btnRightSL = btnVals[1]
			case "btnRightSR":
				btnRightSR = btnVals[1]
			case "btnLeftSL":
				btnLeftSL = btnVals[1]
			case "btnLeftSR":
				btnLeftSR = btnVals[1]
			case "lsLeft":
				lsLeft = btnVals[1]
			case "lsRight":
				lsRight = btnVals[1]
			case "lsUp":
				lsUp = btnVals[1]
			case "lsDown":
				lsDown = btnVals[1]
			case "rsLeft":
				rsLeft = btnVals[1]
			case "rsRight":
				rsRight = btnVals[1]
			case "rsUp":
				rsUp = btnVals[1]
			case "rsDown":
				rsDown = btnVals[1]
		}
	}
	fmt.Println("Key config loaded successfully.")
	return
}
