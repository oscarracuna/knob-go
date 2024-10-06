package main

import (
	"fmt"
	"io"
	"os"

	"github.com/oscarracuna/knob-go/pkg/volume"
)

var (
	Buffer = make([]byte, 24)
)

func main() {
	fmt.Println("Listening...")
	file, err := os.Open("/dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02") // Or you can use the event file as well
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	for {
		_, err := file.Read(Buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading event:", err)
			break
		}
		adjustVolume()
	}
}

func adjustVolume() {
	var EventType = uint16(Buffer[16]) | uint16(Buffer[17])<<8
	var EventCode = uint16(Buffer[18]) | uint16(Buffer[19])<<8
	var EventValue = int32(Buffer[20]) | int32(Buffer[21])<<8 | int32(Buffer[22])<<16 | int32(Buffer[23])<<24

	if EventType == 1 {
		if EventCode == 115 && EventValue == 1 {
			volume.VolumeUp(true)
			fmt.Println("Volume up")
		} else if EventCode == 114 && EventValue == 1 {
			volume.VolumeDown(true)
			fmt.Println("Volume down")
		}
	}

	status := volume.MuteStatus()
	if EventType == 1 {
		if status && EventCode == 113 && EventValue == 1 {
			fmt.Println("Volume muted")
			volume.Mute(true)
		} else if !status && EventCode == 113 && EventValue == 1 {
			fmt.Println("Volume unmuted")
			volume.Unmute(true)
		}
	}
}
