// This contains the required hack for Fell Seal and
// will contain any future hacks for other games.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Joycon js# event node names
var jsDevNameR, jsDevNameL string

// Fell Seal has controller code that uses the controller
// w/o user permission and w/o any way to disable it
// which causes conflicts with the split controller.
// This hack makes the user process unable to access the
// event node while we can still use it with joygo via root.
func fellSealHack () {
	fmt.Printf("\nModding event nodes for Fell Seal...")
	fellMod = true
	// Make sure both joycons are requested & found
	if devPathR == "" || devPathL == "" {
		fmt.Println("\n\nFell Seal requires both joycons to play!")
		fmt.Println("Please set joycons equal to rl or lr in conf.")
		os.Exit(1)
	}
	// Cmd1 gets the path to where the js# node names are
	cmd1R := "ls -R /sys/devices | grep bluetooth | grep " + filepath.Base(devPathR) + " | head -1"
	cmd1L := "ls -R /sys/devices | grep bluetooth | grep " + filepath.Base(devPathL) + " | head -1"
	pathR, cmd1RErr := exec.Command("bash", "-c", cmd1R).Output()
	parseFatal(cmd1RErr, "Cmd1R Fail!")
	pathL, cmd1LErr := exec.Command("bash", "-c", cmd1L).Output()
	parseFatal(cmd1LErr, "Cmd1L Fail!")
	// Cmd2 gets the js# node names
	cmd2R := "ls " + filepath.Dir(string(pathR)) + " | grep js"
	cmd2L := "ls " + filepath.Dir(string(pathL)) + " | grep js"
	jsDevBytesR, cmd2RErr := exec.Command("bash", "-c", cmd2R).Output()
	parseFatal(cmd2RErr, "Cmd2R Fail!")
	jsDevBytesL, cmd2LErr := exec.Command("bash", "-c", cmd2L).Output()
	parseFatal(cmd2LErr, "Cmd2L Fail!")
	jsDevNameR = trimToStr(jsDevBytesR)
	jsDevNameL = trimToStr(jsDevBytesL)
	// Safety check so we don't screw up any of
	// the event node permissions and leave the
	// system in an unwanted/unusable state.
	if jsDevNameR == "" || jsDevNameL == "" {
		fmt.Println("JOYSTICK EVENT NODE(S) NOT FOUND!")
		fmt.Println("ABORTING TO KEEP EVENT DEV PERMISSIONS SAFE!")
		os.Exit(1)
	}
	// Cmd3 turns off group ownership access to the js# nodes
	cmd3R := "chmod 0600 /dev/input/" + jsDevNameR
	cmd3L := "chmod 0600 /dev/input/" + jsDevNameL
	_, cmd3RErr := exec.Command("bash", "-c", cmd3R).Output()
	parseFatal(cmd3RErr, "\nCmd3R Fail! Need root.")
	_, cmd3LErr := exec.Command("bash", "-c", cmd3L).Output()
	parseFatal(cmd3LErr, "Cmd3L Fail!")
	// Cmd4 does the same as cmd3 but to the regular event nodes
	cmd4 := "chmod 0600 "
	_, cmd4RErr := exec.Command("bash", "-c", cmd4 + devPathR).Output()
	parseFatal(cmd4RErr, "Cmd4R Fail!")
	_, cmd4LErr := exec.Command("bash", "-c", cmd4 + devPathL).Output()
	parseFatal(cmd4LErr, "Cmd4L Fail!")
	fmt.Printf("READY!\n")
	return
}

func fellSealCleanUp () {
	fmt.Printf("\nReverting event node permissions...")
	// Additional safety to protect input node integrity
	if jsDevNameR == "" || jsDevNameL == "" {
		fmt.Println("JOYSTICK EVENT NODE NAME(S) NOT FOUND!")
		fmt.Println("ABORTING TO KEEP EVENT DEV PERMISSIONS SAFE!")
		os.Exit(1)
	}
	// Cmd5 reverses cmd3
	cmd5 := "chmod 0660 /dev/input/"
	_, cmd5RErr := exec.Command("bash", "-c", cmd5 + jsDevNameR).Output()
	parseFatal(cmd5RErr, "Cmd5R Fail!")
	_, cmd5LErr := exec.Command("bash", "-c", cmd5 + jsDevNameL).Output()
	parseFatal(cmd5LErr, "Cmd5L Fail!")
	// Cmd6 reverses cmd4
	cmd6 := "chmod 0660 "
	_, cmd6RErr := exec.Command("bash", "-c", cmd6 + devPathR).Output()
	parseFatal(cmd6RErr, "Cmd6R Fail!")
	_, cmd6LErr := exec.Command("bash", "-c", cmd6 + devPathL).Output()
	parseFatal(cmd6LErr, "Cmd6L Fail!")
	fmt.Printf("Done.\n")
}
