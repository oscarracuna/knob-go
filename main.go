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
        exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "+5%").Run()
    } else {
        fmt.Println("Volume down!")
        exec.Command("pactl", "set-sink-volume","@DEFAULT_SINK@", "-5%").Run()
    }
}

func isMute(mute bool) {
  if mute {
    fmt.Println("Volume muted!")
    exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "true").Run()    
  } else {
    fmt.Println("Volume unmuted!")
    exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "false").Run()
  }
}

func main() {
    fmt.Println("Listening...")
    // use your own file here -----
    file, err := os.Open("/dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02")// Or you can use the event file as well  
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

        eventType := uint16(buffer[16]) | uint16(buffer[17])<<8
        eventCode := uint16(buffer[18]) | uint16(buffer[19])<<8
        eventValue := int32(buffer[20]) | int32(buffer[21])<<8 | int32(buffer[22])<<16 | int32(buffer[23])<<24
       
    // These values also work for me but may not work for you
        if eventType == 1 {
            if eventCode == 113 && eventValue == 1{
                isMute(true)
            /* } else if eventCode == 113 && eventValue == 0{
                isMute(false) */
            } else if eventCode == 115 && eventValue == 1 {
                adjustVolume(true)
            } else if eventCode == 114 && eventValue == 1 {
                adjustVolume(false)
            }
        }
    }
}

