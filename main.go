package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
)

func upOrDown(increase bool) {
	if increase {
		fmt.Println("Volume up!")
		exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "+5%").Run()
	} else {
		fmt.Println("Volume down!")
		exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "-5%").Run()
	}
}

func main() {
	fmt.Println("Listening...")
	// use your own id file here -----
	file, err := os.Open("/dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02") // Or you can use the event file as well
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 24) // You may want to change this, too

	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading event:", err)
			break
		}

		var eventType = uint16(buffer[16]) | uint16(buffer[17])<<8
		var eventCode = uint16(buffer[18]) | uint16(buffer[19])<<8
		var eventValue = int32(buffer[20]) | int32(buffer[21])<<8 | int32(buffer[22])<<16 | int32(buffer[23])<<24

		// These values also work for me but may not work for you
		if eventType == 1 {
			if eventCode == 115 && eventValue == 1 {
				upOrDown(true)
			} else if eventCode == 114 && eventValue == 1 {
				upOrDown(false)
			}
			
      // =================================================================
      // temp location of this until I figure out a func or something else
      // =================================================================
      smn := []byte{77, 117, 116, 101, 58, 32, 121, 101, 115, 10}
			nomn := []byte{77, 117, 116, 101, 58, 32, 110, 111, 10}

			out, err := exec.Command("pactl", "get-sink-mute", "@DEFAULT_SINK@").Output()
			if err != nil {
				fmt.Println("Something went wrong with the mute command: ", err)
			}
			si := reflect.DeepEqual(out, smn)
			no := reflect.DeepEqual(out, nomn)
			if eventType == 1 {
				if !si && eventCode == 113 && eventValue == 1{
					fmt.Println("Volume muted!")
					exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "true").Run()
				} else if !no && eventCode == 113 && eventValue == 1{
					fmt.Println("Volume unmuted!")
					exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "false").Run()
				}
			}
		}
	}
}
