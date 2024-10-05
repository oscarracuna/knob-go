package main

import (
    "fmt"
    "os"
    "io"
    "os/exec"
)

func adjustVolume(increase bool) {
    if increase {
        fmt.Println("Volume up!")
        exec.Command("pactl", "set-sink-volume", "1", "+5%").Run()
    } else {
        fmt.Println("Volume down!")
        exec.Command("pactl", "set-sink-volume", "1", "-5%").Run()
    }
}

func main() {
    fmt.Println("Listening...")

    file, err := os.Open("/dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02")
    if err != nil {
        fmt.Println("Error opening event file:", err)
        return
    }
    defer file.Close()

    buffer := make([]byte, 24)

    for {
        _, err := file.Read(buffer)
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error reading event:", err)
            break
        }

        eventType := uint16(buffer[16]) | uint16(buffer[17])<<8
        eventCode := uint16(buffer[18]) | uint16(buffer[19])<<8
        eventValue := int32(buffer[20]) | int32(buffer[21])<<8 | int32(buffer[22])<<16 | int32(buffer[23])<<24

        fmt.Printf("Event type: %d, Event code: %d, Value: %d\n", eventType, eventCode, eventValue)

        if eventType == 1 {
            if eventCode == 115 && eventValue == 1 {
                adjustVolume(true)
            } else if eventCode == 114 && eventValue == 1 {
                adjustVolume(false)
            }
        }
    }
}

