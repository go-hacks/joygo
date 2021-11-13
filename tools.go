// Some basic boiler plate tools
package main

import (
	"fmt"
	"os"
	"github.com/gvalkov/golang-evdev"
)

// Input event matcher
func isEventMatch(event *evdev.InputEvent, matchEvent evMatch) bool {
	if event.Code == matchEvent.Code && event.Type == matchEvent.Type {
		return true
	} else {
		return false
	}
}

// Handle fatal errors
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

// Snip endline char and convert to string
func trimToStr (src []byte) string {
	return string(src[0:len(src)-1])
}
