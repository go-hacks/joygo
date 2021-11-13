// Version 0.80

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/gvalkov/golang-evdev"
)

// Joycon event dev paths
var devPathR, devPathL string
// Joycon in use bools
var rightJoyconIsRequested bool = false
var leftJoyconIsRequested bool = false

// Mod for Fell Seal and possibly other
// games that have the same problem.
// See below for more information on this.
var fellMod bool = false

func main () {
	// Print launch information
	fmt.Println("\nTo disable user level access to joycon")
	fmt.Println("event nodes, run 'sudo ./joygo conf fell'\n")
	fmt.Println("This is necessary for certain games that have")
	fmt.Println("their own conflicting controller handling.\n")

	// Load list of input devices
	devices, err := evdev.ListInputDevices()
	parseFatal(err, "Fail to get input device list")
	// Get dev paths for right joycon if enabled
	if rightJoyconIsRequested {
		for _, dev := range devices {
			if dev.Name == "Joy-Con (R)" {
				fmt.Println("Joy-Con(R) Found! @", dev.Fn)
				devPathR = dev.Fn
				break
			}
		}
	}
	// Get dev paths for left joycon if enabled
	if leftJoyconIsRequested {
		for _, dev := range devices {
			if dev.Name == "Joy-Con (L)" {
				fmt.Println("Joy-Con(L) Found! @", dev.Fn)
				devPathL = dev.Fn
				break
			}
		}
	}

	// Bail if Joycons aren't found
	var failCheck bool = false
	var outStrR string = "Joy-Con(R) Not "
	if devPathR == "" {
		if rightJoyconIsRequested {
			outStrR = outStrR + "Found!"
			failCheck = true
		} else {
			outStrR = outStrR + "Needed."
		}
		fmt.Println(outStrR)
	}
	var outStrL string = "Joy-Con(L) Not "
	if devPathL == "" {
		if leftJoyconIsRequested {
			outStrL = outStrL + "Found!"
			failCheck = true
		} else {
			outStrL = outStrL + "Needed."
		}
		fmt.Println(outStrL)
	}
	if failCheck == true {
		fmt.Println("ABORTING!")
		os.Exit(1)
	}

	// Make event frame channels and initiate handlers
	rightFrameChannel := make(chan []evdev.InputEvent, 5)
	leftFrameChannel := make(chan []evdev.InputEvent, 5)
	if rightJoyconIsRequested {
		go rightJoyconHandler(rightFrameChannel)
	}
	if leftJoyconIsRequested {
		go leftJoyconHandler(leftFrameChannel)
	}

	// Fell Seal hack if requested
	if len(os.Args) == 3 && os.Args[2] == "fell" {
		fellSealHack()
	}

	// Start Ctrl+C & sigterm hook to handle closure
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	exitChan := make(chan bool)
	go exitHook(sigChan, exitChan)

	// Start up the frameReaders
	if rightJoyconIsRequested {
		go frameReader(devPathR, rightFrameChannel)
	}
	if leftJoyconIsRequested {
		go frameReader(devPathL, leftFrameChannel)
	}

	// Wait for exit hook to send exit signal
	fmt.Println("\nCtrl+C to exit...")
	<-exitChan

	return
}

// Reads frames from given device & sends down given channel.
func frameReader (path string, channel chan []evdev.InputEvent) {
	device, _ := evdev.Open(path)
	for {
		eventFrame, _ := device.Read()
		channel <- eventFrame
	}
	return
}

// Wait for Ctrl+C or SIGTERM then cleanup
// and tell main program to quit
func exitHook (sigChan chan os.Signal, exitChan chan bool) {
	<-sigChan
	if fellMod == true {
		fellSealCleanUp()
	}
	exitChan <- true
	return
}
