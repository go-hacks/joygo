package main

import (
  "fmt"
  "os"
  "github.com/gvalkov/golang-evdev"
)

func isEventMatch(event *evdev.InputEvent, matchEvent evMatch) bool {
	if event.Code == matchEvent.Code && event.Type == matchEvent.Type && event.Value == matchEvent.Value {
		return true
	} else {
	  return false
  }
}

func parseFatal(err error, msg string) {
	if err != nil {
		if msg != "" {
			fmt.Println(msg)
		}
		fmt.Println(err)
		os.Exit(1)
	} else {
		return
	}
}
