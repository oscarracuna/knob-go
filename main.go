package main 

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("ola")
  //This is for my specific keeb
  //Went for symlink instead of event file
  file, err := os.Open("/dev/input/by-id/usb-Telink_Wireless_Gaming_Keyboard-event-kbd")
  if err != nil {
    fmt.Println("Something went wrong with the file: ", err)
    return
  }
  defer file.Close()
}

// Event thingy
// /dev/input/by-id/usb-Telink_Wireless_Gaming_Keyboard-event-kbd

