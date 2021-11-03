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
var btnXPress = evMatch{
	Code: 305, Type: 1, Value: 1,
}
var btnXRelease = evMatch{
	Code: 305, Type: 1, Value: 0,
}
var btnXHold = evMatch{
	Code: 305, Type: 1, Value: 2,
}
var btnYPress = evMatch{
	Code: 307, Type: 1, Value: 1,
}
var btnYRelease = evMatch{
	Code: 307, Type: 1, Value: 0,
}
var btnYHold = evMatch{
	Code: 307, Type: 1, Value: 2,
}
var btnAPress = evMatch{
	Code: 304, Type: 1, Value: 1,
}
var btnARelease = evMatch{
	Code: 304, Type: 1, Value: 0,
}
var btnAHold = evMatch{
	Code: 304, Type: 1, Value: 2,
}
var btnBPress = evMatch{
	Code: 306, Type: 1, Value: 1,
}
var btnBRelease = evMatch{
	Code: 306, Type: 1, Value: 0,
}
var btnBHold = evMatch{
	Code: 306, Type: 1, Value: 2,
}
var btnPPress = evMatch{
	Code: 313, Type: 1, Value: 1,
}
var btnPRelease = evMatch{
	Code: 313, Type: 1, Value: 0,
}
var btnPHold = evMatch{
	Code: 313, Type: 1, Value: 2,
}
var btnHPress = evMatch{
	Code: 316, Type: 1, Value: 1,
}
var btnHRelease = evMatch{
	Code: 316, Type: 1, Value: 0,
}
var btnHHold = evMatch{
	Code: 316, Type: 1, Value: 2,
}
var btnRPress = evMatch{
	Code: 318, Type: 1, Value: 1,
}
var btnRRelease = evMatch{
	Code: 318, Type: 1, Value: 0,
}
var btnRHold = evMatch{
	Code: 318, Type: 1, Value: 2,
}
var btnRZPress = evMatch{
	Code: 319, Type: 1, Value: 1,
}
var btnRZRelease = evMatch{
	Code: 319, Type: 1, Value: 0,
}
var btnRZHold = evMatch{
	Code: 319, Type: 1, Value: 2,
}
var btnSLPress = evMatch{
	Code: 308, Type: 1, Value: 1,
}
var btnSLRelease = evMatch{
	Code: 308, Type: 1, Value: 0,
}
var btnSLHold = evMatch{
	Code: 308, Type: 1, Value: 2,
}
var btnSRPress = evMatch{
	Code: 309, Type: 1, Value: 1,
}
var btnSRRelease = evMatch{
	Code: 309, Type: 1, Value: 0,
}
var btnSRHold = evMatch{
	Code: 309, Type: 1, Value: 2,
}
var rightStickHLeft = evMatch{
  Code: 17, Type: 3, Value: -1,
}
var rightStickHRight = evMatch{
  Code: 17, Type: 3, Value: 1,
}
var rightStickHRelease = evMatch{
  Code: 17, Type: 3, Value: 0,
}
var rightStickVDown = evMatch{
  Code: 16, Type: 3, Value: -1,
}
var rightStickVUp = evMatch{
  Code: 16, Type: 3, Value: 1,
}
var rightStickVRelease = evMatch{
  Code: 16, Type: 3, Value: 0,
}
var btnR3Press = evMatch{
  Code: 315, Type: 1, Value: 1,
}
var btnR3Release = evMatch{
  Code: 315, Type: 1, Value: 0,
}
var btnR3Hold = evMatch{
  Code: 315, Type: 1, Value: 2,
}
// Left Joycon events
var btnDownPress = evMatch{
	Code: 305, Type: 1, Value: 1,
}
var btnDownRelease = evMatch{
	Code: 305, Type: 1, Value: 0,
}
var btnDownHold = evMatch{
	Code: 305, Type: 1,	Value: 2,
}
var btnRightPress = evMatch{
	Code: 307, Type: 1,	Value: 1,
}
var btnRightRelease = evMatch{
	Code: 307, Type: 1,	Value: 0,
}
var btnRightHold = evMatch{
	Code: 307, Type: 1,	Value: 2,
}
var btnLeftPress = evMatch{
	Code: 304, Type: 1,	Value: 1,
}
var btnLeftRelease = evMatch{
	Code: 304, Type: 1,	Value: 0,
}
var btnLeftHold = evMatch{
	Code: 304, Type: 1,	Value: 2,
}
var btnUpPress = evMatch{
	Code: 306, Type: 1,	Value: 1,
}
var btnUpRelease = evMatch{
	Code: 306, Type: 1,	Value: 0,
}
var btnUpHold = evMatch{
	Code: 306, Type: 1,	Value: 2,
}
var btnMPress = evMatch{
	Code: 312, Type: 1,	Value: 1,
}
var btnMRelease = evMatch{
	Code: 312, Type: 1,	Value: 0,
}
var btnMHold = evMatch{
	Code: 312, Type: 1,	Value: 2,
}
var btnSSPress = evMatch{
	Code: 317, Type: 1,	Value: 1,
}
var btnSSRelease = evMatch{
	Code: 317, Type: 1,	Value: 0,
}
var btnSSHold = evMatch{
	Code: 317, Type: 1,	Value: 2,
}
var btnLPress = evMatch{
	Code: 318, Type: 1,	Value: 1,
}
var btnLRelease = evMatch{
	Code: 318, Type: 1,	Value: 0,
}
var btnLHold = evMatch{
	Code: 318, Type: 1, Value: 2,
}
var btnLZPress = evMatch{
	Code: 319, Type: 1,	Value: 1,
}
var btnLZRelease = evMatch{
	Code: 319, Type: 1, Value: 0,
}
var btnLZHold = evMatch{
	Code: 319, Type: 1,	Value: 2,
}
var leftStickHRight = evMatch{
  Code: 17, Type: 3, Value: -1,
}
var leftStickHLeft = evMatch{
  Code: 17, Type: 3, Value: 1,
}
var leftStickHRelease = evMatch{
  Code: 17, Type: 3, Value: 0,
}
var leftStickVUp = evMatch{
  Code: 16, Type: 3, Value: -1,
}
var leftStickVDown = evMatch{
  Code: 16, Type: 3, Value: 1,
}
var leftStickVRelease = evMatch{
  Code: 16, Type: 3, Value: 0,
}
var btnL3Press = evMatch{
  Code: 314, Type: 1, Value: 1,
}
var btnL3Release = evMatch{
  Code: 314, Type: 1, Value: 0,
}
var btnL3Hold = evMatch{
  Code: 314, Type: 1, Value: 2,
}
// Other events
var endFrame = evMatch{
  Code: 0, Type: 0, Value: 0,
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
