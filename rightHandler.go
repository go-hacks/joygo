package main

import (
  "fmt"
  "github.com/gvalkov/golang-evdev"
  "github.com/go-vgo/robotgo"
)

func rightJoyconHandler (frameChan chan []evdev.InputEvent) {
  var rightStickHState int32
  var rightStickVState int32
  for {
    // Get event frame from channel
    eventFrame := <-frameChan
    for _, event := range eventFrame {
      if isEventMatch(&event, endFrame) {
        continue
      }
      if isEventMatch(&event, btnXPress) {
        robotgo.KeyToggle(btnX, "down")
        fmt.Println("X Press")
      } else if isEventMatch(&event, btnXRelease) {
        robotgo.KeyToggle(btnX, "up")
        fmt.Println("X Release")
      } else if isEventMatch(&event, btnXHold) {
        fmt.Println("X Hold")
      } else if isEventMatch(&event, btnYPress) {
        robotgo.KeyToggle(btnY, "down")
        fmt.Println("Y Press")
      } else if isEventMatch(&event, btnYRelease) {
        robotgo.KeyToggle(btnY, "up")
        fmt.Println("Y Release")
      } else if isEventMatch(&event, btnYHold) {
        fmt.Println("Y Hold")
      } else if isEventMatch(&event, btnAPress) {
        robotgo.KeyToggle(btnA, "down")
        fmt.Println("A Press")
      } else if isEventMatch(&event, btnARelease) {
        robotgo.KeyToggle(btnA, "up")
        fmt.Println("A Release")
      } else if isEventMatch(&event, btnAHold) {
        fmt.Println("A Hold")
      } else if isEventMatch(&event, btnBPress) {
        robotgo.KeyToggle(btnB, "down")
        fmt.Println("B Press")
      } else if isEventMatch(&event, btnBRelease) {
        robotgo.KeyToggle(btnB, "up")
        fmt.Println("B Release")
      } else if isEventMatch(&event, btnBHold) {
        fmt.Println("B Hold")
      } else if isEventMatch(&event, btnPPress) {
        robotgo.KeyToggle(btnPlus, "down")
        fmt.Println("+ Press")
      } else if isEventMatch(&event, btnPRelease) {
        robotgo.KeyToggle(btnPlus, "up")
        fmt.Println("+ Release")
      } else if isEventMatch(&event, btnPHold) {
        fmt.Println("+ Hold")
      } else if isEventMatch(&event, btnHPress) {
        robotgo.KeyToggle(btnHome, "down")
        fmt.Println("Home Press")
      } else if isEventMatch(&event, btnHRelease) {
        robotgo.KeyToggle(btnHome, "up")
        fmt.Println("Home Release")
      } else if isEventMatch(&event, btnHHold) {
        fmt.Println("Home Hold")
      } else if isEventMatch(&event, btnRPress) {
        robotgo.KeyToggle(btnR, "down")
        fmt.Println("R Press")
      } else if isEventMatch(&event, btnRRelease) {
        robotgo.KeyToggle(btnR, "up")
        fmt.Println("R Release")
      } else if isEventMatch(&event, btnRHold) {
        fmt.Println("R Hold")
      } else if isEventMatch(&event, btnRZPress) {
        robotgo.KeyToggle(btnRZ, "down")
        fmt.Println("RZ Press")
      } else if isEventMatch(&event, btnRZRelease) {
        robotgo.KeyToggle(btnRZ, "up")
        fmt.Println("RZ Release")
      } else if isEventMatch(&event, btnRZHold) {
        fmt.Println("RZ Hold")
      } else if isEventMatch(&event, btnSLPress) {
        robotgo.KeyToggle(btnRightSL, "down")
        fmt.Println("SL Press")
      } else if isEventMatch(&event, btnSLRelease) {
        robotgo.KeyToggle(btnRightSL, "up")
        fmt.Println("SL Release")
      } else if isEventMatch(&event, btnSLHold) {
        fmt.Println("SL Hold")
      } else if isEventMatch(&event, btnSRPress) {
        robotgo.KeyToggle(btnRightSR, "down")
        fmt.Println("SR Press")
      } else if isEventMatch(&event, btnSRRelease) {
        robotgo.KeyToggle(btnRightSR, "up")
        fmt.Println("SR Release")
      } else if isEventMatch(&event, btnSRHold) {
        fmt.Println("SR Hold")
      } else if isEventMatch(&event, rightStickHLeft) {
        rightStickHState = event.Value
        robotgo.KeyToggle(rsLeft, "down")
        fmt.Println("Right Stick H Left")
      } else if isEventMatch(&event, rightStickHRight) {
        rightStickHState = event.Value
        robotgo.KeyToggle(rsRight, "down")
        fmt.Println("Right Stick H Right")
      } else if isEventMatch(&event, rightStickHRelease) {
        switch rightStickHState {
        case 1:
          robotgo.KeyToggle(rsRight, "up")
        case -1:
          robotgo.KeyToggle(rsLeft, "up")
        }
        rightStickHState = event.Value
        fmt.Println("Right Stick H Release")
      } else if isEventMatch(&event, rightStickVUp) {
        rightStickVState = event.Value
        robotgo.KeyToggle(rsUp, "down")
        fmt.Println("Right Stick V Up")
      } else if isEventMatch(&event, rightStickVDown) {
        rightStickVState = event.Value
        robotgo.KeyToggle(rsDown, "down")
        fmt.Println("Right Stick V Down")
      } else if isEventMatch(&event, rightStickVRelease) {
        switch rightStickVState {
        case 1:
          robotgo.KeyToggle(rsUp, "up")
        case -1:
          robotgo.KeyToggle(rsDown, "up")
        }
        rightStickVState = event.Value
        fmt.Println("Right Stick V Release")
      } else if isEventMatch(&event, btnR3Press) {
        robotgo.KeyToggle(btnR3, "down")
        fmt.Println("R3 Press")
      } else if isEventMatch(&event, btnR3Release) {
        robotgo.KeyToggle(btnR3, "up")
        fmt.Println("R3 Release")
      } else if isEventMatch(&event, btnR3Hold) {
        fmt.Println("R3 Hold")
      }
    }
  }
  return
}
