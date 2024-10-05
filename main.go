package main

import (
    "fmt"
    "os"
    "io"
    "os/exec"
    "encoding/binary"
)

// /dev/input/by-id/usb-RDR_EPOMAKER_Shadow-S-event-if02

// amixer did not work
// pactl did not work either
// Figure out if issue is the exec.Commands or something else
func adjustVolume(increase bool) {
    //var cmd *exec.Cmd

    if increase {
        exec.Command("pactl", "set-sink-volume", "1", "+5%").Run()
    } else {
        exec.Command("pactl", "set-sink-volume", "1", "-5%").Run()
    }

    //err := cmd.Run()
    //if err != nil {
    //  fmt.Println("Error with command: ", err)
    //}
}

func main() {
    fmt.Println("Listening...")

    file, err := os.Open("/dev/input/event16")
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

    fmt.Printf("Buffer content: %v\n", buffer)
    fmt.Printf("Event type: %d, Event code: %d, Value: %d\n", buffer[18], buffer[20], buffer[22])

    // got this from chatgpt idk

    
    type InputEvent struct {
    Time  [8]byte   // First 16 bytes: Time (sec, usec)
    Type  uint16    // Event type (2 bytes)
    Code  uint16    // Event code (2 bytes)
    Value int32     // Event value (4 bytes)
    }

    var event InputEvent
    for {
    err := binary.Read(file, binary.LittleEndian, &event)
    if err != nil {
        fmt.Println("Error reading event:", err)
        break
      }

    fmt.Printf("Event type: %d, Event code: %d, Value: %d\n", event.Type, event.Code, event.Value)
    }
    
    if eventType == 1 {
      if eventCode == 115 && value == 1 {
        adjustVolume(true)
      } else if eventCode == 114 && value == 1 {
        adjustVolume(false)
      }
    }    
  }
}

