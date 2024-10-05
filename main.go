package main

import (
    "fmt"
    "os"
    "io"
    "os/exec"
)

// /dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02

// amixer did not work
// pactl did not work either
// Figure out if issue is the exec.Commands or something else
func adjustVolume(increase bool) {
    if increase {
        exec.Command("pactl", "set-sink-volume", "1", "+5%").Run()
    } else {
        exec.Command("pactl", "set-sink-volume", "1", "-5%").Run()
    }
}

func main() {
    fmt.Println("Listening...")

    file, err := os.Open("/dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02")
    if err != nil {
        fmt.Println("Something went wrong with the file: ", err)
        return
    }
    defer file.Close()

    buffer := make([]byte, 24)

    for {
        _, err := file.Read(buffer)
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error reading event: ", err)
            break
        }
    eventType := buffer[18]
    eventCode := buffer[20]
    value := buffer[22]

    if eventType == 1 {
      if eventCode == 115 && value == 1 {
        adjustVolume(true)
      } else if eventCode == 114 && value == 1 {
        adjustVolume(false)
      }
    }
  }
}

