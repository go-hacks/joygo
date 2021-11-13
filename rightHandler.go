// Function for interpretting the input
// events of the right Joycon controller
package main

import (
	"fmt"
	"github.com/gvalkov/golang-evdev"
	"github.com/go-vgo/robotgo"
)

func rightJoyconHandler (frameChan chan []evdev.InputEvent) {
	var rightStickHState int32
	var rightStickVState int32
	var btnStr string
	var isBtn bool
	for {
		// Get event frame from channel
		eventFrame := <-frameChan
		// Parse events in frame
		for _, event := range eventFrame {
			if isEventMatch(&event, endFrameEvent) {
				continue
			}
			if isEventMatch(&event, btnXEvent) {
				btnStr = "X"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnX, "up")
					case 1:
						robotgo.KeyToggle(btnX, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnYEvent) {
				btnStr = "Y"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnY, "up")
					case 1:
						robotgo.KeyToggle(btnY, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnAEvent) {
				btnStr = "A"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnA, "up")
					case 1:
						robotgo.KeyToggle(btnA, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnBEvent) {
				btnStr = "B"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnB, "up")
					case 1:
						robotgo.KeyToggle(btnB, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnPlusEvent) {
				btnStr = "+"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnPlus, "up")
					case 1:
						robotgo.KeyToggle(btnPlus, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnHomeEvent) {
				btnStr = "Home"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnHome, "up")
					case 1:
						robotgo.KeyToggle(btnHome, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnREvent) {
				btnStr = "R"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnR, "up")
					case 1:
						robotgo.KeyToggle(btnR, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnRZEvent) {
				btnStr = "RZ"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnRZ, "up")
					case 1:
						robotgo.KeyToggle(btnRZ, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnSLEvent) {
				btnStr = "Right SL"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnRightSL, "up")
					case 1:
						robotgo.KeyToggle(btnRightSL, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnSREvent) {
				btnStr = "Right SR"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnRightSR, "up")
					case 1:
						robotgo.KeyToggle(btnRightSR, "down")
					case 2:
				}
			} else if isEventMatch(&event, rightStickHEvent) {
				isBtn = false
				switch event.Value {
					case 1:
						btnStr = "Right Stick H Right"
						robotgo.KeyToggle(rsRight, "down")
					case 0:
						btnStr = "Right Stick H Release"
						switch rightStickHState {
							case 1:
								robotgo.KeyToggle(rsRight, "up")
							case -1:
								robotgo.KeyToggle(rsLeft, "up")
						}
					case -1:
						btnStr = "Right Stick H Left"
						robotgo.KeyToggle(rsLeft, "down")
				}
				rightStickHState = event.Value
			} else if isEventMatch(&event, rightStickVEvent) {
				isBtn = false
				switch event.Value {
					case 1:
						btnStr = "Right Stick V Up"
						robotgo.KeyToggle(rsUp, "down")
					case 0:
						btnStr = "Right Stick V Release"
						switch rightStickVState {
							case 1:
								robotgo.KeyToggle(rsUp, "up")
							case -1:
								robotgo.KeyToggle(rsDown, "up")
						}
					case -1:
						btnStr = "Right Stick V Down"
						robotgo.KeyToggle(rsDown, "down")
				}
				rightStickVState = event.Value
			} else if isEventMatch(&event, btnR3Event) {
				btnStr = "R3"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnR3, "up")
					case 1:
						robotgo.KeyToggle(btnR3, "down")
					case 2:
				}
			}
			switch isBtn {
				case true:
					switch event.Value {
						case 1:
							fmt.Printf("%s Press\n", btnStr)
						case 0:
							fmt.Printf("%s Release\n", btnStr)
						case 2:
							fmt.Printf("%s Hold\n", btnStr)
					}
				// For joysticks
				case false:
					fmt.Printf("%s\n", btnStr)
			}
		}
	}
	return
}
