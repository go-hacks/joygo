// Function for interpretting the input
// events of the left Joycon controller
package main

import (
	"fmt"
	"github.com/gvalkov/golang-evdev"
	"github.com/go-vgo/robotgo"
)

func leftJoyconHandler (frameChan chan []evdev.InputEvent) {
	var leftStickHState int32
	var leftStickVState int32
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
			if isEventMatch(&event, btnDownEvent) {
				btnStr = "Down"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnDown, "up")
					case 1:
						robotgo.KeyToggle(btnDown, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnRightEvent) {
				btnStr = "Right"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnRight, "up")
					case 1:
						robotgo.KeyToggle(btnRight, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnLeftEvent) {
				btnStr = "Left"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnLeft, "up")
					case 1:
						robotgo.KeyToggle(btnLeft, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnUpEvent) {
				btnStr = "Up"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnUp, "up")
					case 1:
						robotgo.KeyToggle(btnUp, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnMinusEvent) {
				btnStr = "-"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnMinus, "up")
					case 1:
						robotgo.KeyToggle(btnMinus, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnSSEvent) {
				btnStr = "SS"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnScreenShot, "up")
					case 1:
						robotgo.KeyToggle(btnScreenShot, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnLEvent) {
				btnStr = "L"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnL, "up")
					case 1:
						robotgo.KeyToggle(btnL, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnLZEvent) {
				btnStr = "LZ"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnLZ, "up")
					case 1:
						robotgo.KeyToggle(btnLZ, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnSLEvent) {
				btnStr = "Left SL"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnLeftSL, "up")
					case 1:
						robotgo.KeyToggle(btnLeftSL, "down")
					case 2:
				}
			} else if isEventMatch(&event, btnSREvent) {
				btnStr = "Left SR"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnLeftSR, "up")
					case 1:
						robotgo.KeyToggle(btnLeftSR, "down")
					case 2:
				}
			} else if isEventMatch(&event, leftStickHEvent) {
				isBtn = false
				switch event.Value {
					case 1:
						btnStr = "Left Stick H Left"
						robotgo.KeyToggle(lsLeft, "down")
					case 0:
						btnStr = "Left Stick H Release"
						switch leftStickHState {
							case 1:
								robotgo.KeyToggle(lsLeft, "up")
							case -1:
								robotgo.KeyToggle(lsRight, "up")
						}
					case -1:
						btnStr = "Left Stick H Right"
						robotgo.KeyToggle(lsRight, "down")
				}
				leftStickHState = event.Value
			} else if isEventMatch(&event, leftStickVEvent) {
				isBtn = false
				switch event.Value {
					case 1:
						btnStr = "Left Stick V Down"
						robotgo.KeyToggle(lsDown, "down")
					case 0:
						btnStr = "Left Stick V Release"
						switch leftStickVState {
							case 1:
								robotgo.KeyToggle(lsDown, "up")
							case -1:
								robotgo.KeyToggle(lsUp, "up")
						}
					case -1:
						btnStr = "Left Stick V Up"
						robotgo.KeyToggle(lsUp, "down")
				}
				leftStickVState = event.Value
			} else if isEventMatch(&event, btnL3Event) {
				btnStr = "L3"
				isBtn = true
				switch event.Value {
					case 0:
						robotgo.KeyToggle(btnL3, "up")
					case 1:
						robotgo.KeyToggle(btnL3, "down")
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
