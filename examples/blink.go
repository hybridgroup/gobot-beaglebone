package main

import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot-beaglebone"
  "time"
)

func main() {

  beaglebone := new(gobotBeaglebone.Beaglebone)
  beaglebone.Name = "Beaglebone"

  led := gobotBeaglebone.NewLed(beaglebone)
  led.Driver = gobot.Driver{
    Name: "led",
    Pin: "P9_12",
  }

  connections := []interface{} {
    beaglebone,
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
