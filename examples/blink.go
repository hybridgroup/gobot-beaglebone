package main

import (
  "github.com/hybridgroup/gobot/src/gobot"
  "github.com/hybridgroup/gobot/beaglebone"
  "time"
)

func main() {

  bone := new(beaglebone.Beaglebone)
  bone.Name = "Beaglebone"

  led := beaglebone.NewLed(bone)
  led.Driver = gobot.Driver{
    Name: "led",
    Pin: "P9_12",
  }

  connections := []interface{} {
    bone,
  }
  devices := []interface{} {
    led,
  }

  work := func(){
    gobot.Every(1000 * time.Millisecond, func(){ led.Toggle() })
  }
  
  robot := gobot.Robot{
      Connections: connections, 
      Devices: devices,
      Work: work,
  }

  robot.Start()
}
