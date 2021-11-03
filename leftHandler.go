package main

import (
  "fmt"
  "github.com/gvalkov/golang-evdev"
  "github.com/go-vgo/robotgo"
)

func leftJoyconHandler (frameChan chan []evdev.InputEvent) {
  var leftStickHState int32
  var leftStickVState int32
  for {
    // Get event frame from channel
    eventFrame := <-frameChan
    for _, event := range eventFrame {
      if isEventMatch(&event, endFrame) {
        continue
      }
      if isEventMatch(&event, btnDownPress) {
        robotgo.KeyToggle(btnDown, "down")
        fmt.Println("Down Press")
      } else if isEventMatch(&event, btnDownRelease) {
        robotgo.KeyToggle(btnDown, "up")
        fmt.Println("Down Release")
      } else if isEventMatch(&event, btnDownHold) {
        fmt.Println("Down Hold")
      } else if isEventMatch(&event, btnRightPress) {
        robotgo.KeyToggle(btnRight, "down")
        fmt.Println("Right Press")
      } else if isEventMatch(&event, btnRightRelease) {
        robotgo.KeyToggle(btnRight, "up")
        fmt.Println("Right Release")
      } else if isEventMatch(&event, btnRightHold) {
        fmt.Println("Right Hold")
      } else if isEventMatch(&event, btnLeftPress) {
        robotgo.KeyToggle(btnLeft, "down")
        fmt.Println("Left Press")
      } else if isEventMatch(&event, btnLeftRelease) {
        robotgo.KeyToggle(btnLeft, "up")
        fmt.Println("Left Release")
      } else if isEventMatch(&event, btnLeftHold) {
        fmt.Println("Left Hold")
      } else if isEventMatch(&event, btnUpPress) {
        robotgo.KeyToggle(btnUp, "down")
        fmt.Println("Up Press")
      } else if isEventMatch(&event, btnUpRelease) {
        robotgo.KeyToggle(btnUp, "up")
        fmt.Println("Up Release")
      } else if isEventMatch(&event, btnUpHold) {
        fmt.Println("Up Hold")
      } else if isEventMatch(&event, btnMPress) {
        robotgo.KeyToggle(btnMinus, "down")
        fmt.Println("- Press")
      } else if isEventMatch(&event, btnMRelease) {
        robotgo.KeyToggle(btnMinus, "up")
        fmt.Println("- Release")
      } else if isEventMatch(&event, btnMHold) {
        fmt.Println("- Hold")
      } else if isEventMatch(&event, btnSSPress) {
        robotgo.KeyToggle(btnScreenShot, "down")
        fmt.Println("SS Press")
      } else if isEventMatch(&event, btnSSRelease) {
        robotgo.KeyToggle(btnScreenShot, "up")
        fmt.Println("SS Release")
      } else if isEventMatch(&event, btnSSHold) {
        fmt.Println("SS Hold")
      } else if isEventMatch(&event, btnLPress) {
        robotgo.KeyToggle(btnL, "down")
        fmt.Println("L Press")
      } else if isEventMatch(&event, btnLRelease) {
        robotgo.KeyToggle(btnL, "up")
        fmt.Println("L Release")
      } else if isEventMatch(&event, btnLHold) {
        fmt.Println("L Hold")
      } else if isEventMatch(&event, btnLZPress) {
        robotgo.KeyToggle(btnLZ, "down")
        fmt.Println("LZ Press")
      } else if isEventMatch(&event, btnLZRelease) {
        robotgo.KeyToggle(btnLZ, "up")
        fmt.Println("LZ Release")
      } else if isEventMatch(&event, btnLZHold) {
        fmt.Println("LZ Hold")
      } else if isEventMatch(&event, btnSLPress) {
        robotgo.KeyToggle(btnLeftSL, "down")
        fmt.Println("SL Press")
      } else if isEventMatch(&event, btnSLRelease) {
        robotgo.KeyToggle(btnLeftSL, "up")
        fmt.Println("SL Release")
      } else if isEventMatch(&event, btnSLHold) {
        fmt.Println("SL Hold")
      } else if isEventMatch(&event, btnSRPress) {
        robotgo.KeyToggle(btnLeftSR, "down")
        fmt.Println("SR Press")
      } else if isEventMatch(&event, btnSRRelease) {
        robotgo.KeyToggle(btnLeftSR, "up")
        fmt.Println("SR Release")
      } else if isEventMatch(&event, btnSRHold) {
        fmt.Println("SR Hold")
      } else if isEventMatch(&event, leftStickHLeft) {
        leftStickHState = event.Value
        robotgo.KeyToggle(lsLeft, "down")
        fmt.Println("Left Stick H Left")
      } else if isEventMatch(&event, leftStickHRight) {
        leftStickHState = event.Value
        robotgo.KeyToggle(lsRight, "down")
        fmt.Println("Left Stick H Right")
      } else if isEventMatch(&event, leftStickHRelease) {
        switch leftStickHState {
        case 1:
          robotgo.KeyToggle(lsLeft, "up")
        case -1:
          robotgo.KeyToggle(lsRight, "up")
        }
        leftStickHState = event.Value
        fmt.Println("Left Stick H Release")
      } else if isEventMatch(&event, leftStickVUp) {
        leftStickVState = event.Value
        robotgo.KeyToggle(lsUp, "down")
        fmt.Println("Left Stick V Up")
      } else if isEventMatch(&event, leftStickVDown) {
        leftStickVState = event.Value
        robotgo.KeyToggle(lsDown, "down")
        fmt.Println("Left Stick V Down")
      } else if isEventMatch(&event, leftStickVRelease) {
        switch leftStickVState {
        case 1:
          robotgo.KeyToggle(lsDown, "up")
        case -1:
          robotgo.KeyToggle(lsUp, "up")
        }
        leftStickVState = event.Value
        fmt.Println("Left Stick V Release")
      } else if isEventMatch(&event, btnL3Press) {
        robotgo.KeyToggle(btnL3, "down")
        fmt.Println("L3 Press")
      } else if isEventMatch(&event, btnL3Release) {
        robotgo.KeyToggle(btnL3, "up")
        fmt.Println("L3 Release")
      } else if isEventMatch(&event, btnL3Hold) {
        fmt.Println("L3 Hold")
      }
    }
  }
  return
}
