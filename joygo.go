// Version 0.77

package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"github.com/gvalkov/golang-evdev"
)

// Joycon in use bools
var joyconR, joyconL bool = false, false

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
	var devPathR string
	if joyconR {
		for _, dev := range devices {
			if dev.Name == "Joy-Con (R)" {
				fmt.Println("Joy-Con(R) Found! @", dev.Fn)
				devPathR = dev.Fn
				break
			}
		}
	}
	// Get dev paths for left joycon if enabled
	var devPathL string
	if joyconL {
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
		if joyconR {
			outStrR = outStrR + "Found!"
			failCheck = true
		} else {
			outStrR = outStrR + "Needed."
		}
		fmt.Println(outStrR)
	}
	var outStrL string = "Joy-Con(L) Not "
	if devPathL == "" {
		if joyconL {
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
	if joyconR {
		go rightJoyconHandler(rightFrameChannel)
	}
	if joyconL {
		go leftJoyconHandler(leftFrameChannel)
	}

	// Fell Seal has controller code that uses the controller
	// w/o user permission and w/o any way to disable it
	// which causes conflicts with the split controller.
	// This hack makes the user process unable to access the
	// event node while we can still use it with joygo via root.
	var jsDevNameR string
	var jsDevNameL string
	if len(os.Args) == 3 && os.Args[2] == "fell" {
		fmt.Printf("Modding event nodes for Fell Seal...")
		fellMod = true
		evNumR := strings.Split(devPathR, "/")
		evNumL := strings.Split(devPathL, "/")
		cmd1R := "ls -R /sys/devices | grep bluetooth | grep " + evNumR[len(evNumR)-1] + " | head -1"
		cmd1L := "ls -R /sys/devices | grep bluetooth | grep " + evNumL[len(evNumL)-1] + " | head -1"
		pathR, cmd1RErr := exec.Command("bash", "-c", cmd1R).Output()
		parseFatal(cmd1RErr, "Cmd1R Fail!")
		pathL, cmd1LErr := exec.Command("bash", "-c", cmd1L).Output()
		parseFatal(cmd1LErr, "Cmd1L Fail!")
		splitPathR := strings.Split(string(pathR), "/")
		splitPathL := strings.Split(string(pathL), "/")
		var newPathR string
		for i := 1; i < len(splitPathR)-1; i++ {
			newPathR = newPathR + "/" + splitPathR[i]
		}
		var newPathL string
		for i := 1; i < len(splitPathL)-1; i++ {
			newPathL = newPathL + "/" + splitPathL[i]
		}
		cmd2R := "ls " + newPathR + " | grep js"
		cmd2L := "ls " + newPathL + " | grep js"
		jsDevBytesR, cmd2RErr := exec.Command("bash", "-c", cmd2R).Output()
		parseFatal(cmd2RErr, "Cmd2R Fail!")
		jsDevBytesL, cmd2LErr := exec.Command("bash", "-c", cmd2L).Output()
		parseFatal(cmd2LErr, "Cmd2L Fail!")
		jsDevNameR = string(jsDevBytesR[0:len(jsDevBytesR)-1])
		jsDevNameL = string(jsDevBytesL[0:len(jsDevBytesL)-1])
		cmd3 := "chmod 0600 /dev/input/"
		_, cmd3RErr := exec.Command("bash", "-c", cmd3 + jsDevNameR).Output()
		parseFatal(cmd3RErr, "\nCmd3R Fail! Need root.")
		_, cmd3LErr := exec.Command("bash", "-c", cmd3 + jsDevNameL).Output()
		parseFatal(cmd3LErr, "Cmd3L Fail!")
		cmd4 := "chmod 0600 "
		_, cmd4RErr := exec.Command("bash", "-c", cmd4 + devPathR).Output()
		parseFatal(cmd4RErr, "Cmd4R Fail!")
		_, cmd4LErr := exec.Command("bash", "-c", cmd4 + devPathL).Output()
		parseFatal(cmd4LErr, "Cmd4L Fail!")
		fmt.Printf("READY!\n")
	}

	// Start Ctrl+C & sigterm hook to handle closure
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		if fellMod == true {
			fmt.Printf("Reverting event node permissions...")
			cmd5 := "chmod 0660 /dev/input/"
			_, cmd5RErr := exec.Command("bash", "-c", cmd5 + jsDevNameR).Output()
			parseFatal(cmd5RErr, "Cmd5R Fail!")
			_, cmd5LErr := exec.Command("bash", "-c", cmd5 + jsDevNameL).Output()
			parseFatal(cmd5LErr, "Cmd5L Fail!")
			cmd6 := "chmod 0660 "
			_, cmd6RErr := exec.Command("bash", "-c", cmd6 + devPathR).Output()
			parseFatal(cmd6RErr, "Cmd6R Fail!")
			_, cmd6LErr := exec.Command("bash", "-c", cmd6 + devPathL).Output()
			parseFatal(cmd6LErr, "Cmd6L Fail!")
			fmt.Printf("Done.\n")
		}
		os.Exit(0)
	}()

	// Start up the frameReaders
	if joyconR {
		go frameReader(devPathR, rightFrameChannel)
	}
	if joyconL {
		go frameReader(devPathL, leftFrameChannel)
	}

	// Main program just waits for interrupt/sigterm.
	// I put a ticker here so it's not sitting in an
	// unlimited endless loop going brrrrrrrrrrrr XD
	fmt.Println("\nCtrl+C to exit...")
	ticker := time.NewTicker(10 * time.Second).C
	for {
		select {
			case <-ticker:
		}
	}
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
