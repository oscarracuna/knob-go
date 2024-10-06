package volume

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
)

var (
	si_bytes = []byte{77, 117, 116, 101, 58, 32, 121, 101, 115, 10}
	no_bytes = []byte{77, 117, 116, 101, 58, 32, 110, 111, 10}
	status   bool
)

func VolumeUp(up bool) {
	if up {
		err := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "+5%").Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func VolumeDown(down bool) {
	if down {
		err := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "-5%").Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Mute(mute bool) {
	if mute {
		err := exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "true").Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Unmute(unmute bool) {
	if unmute {
		err := exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "false").Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func MuteStatus() bool {
	mute_cmd, err := exec.Command("pactl", "get-sink-mute", "@DEFAULT_SINK@").Output()
	if err != nil {
		fmt.Println("Something went wrong with the mute command: ", err)
	}
	si := reflect.DeepEqual(mute_cmd, si_bytes)
	no := reflect.DeepEqual(mute_cmd, no_bytes)

	if !si {
		status = true

	}
	if !no {
		status = false

	}
	return status
}
