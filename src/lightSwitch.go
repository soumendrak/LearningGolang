package main

import "fmt"

var lightSwitchIsOn bool = false

func main() {
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
}

func printLightSwitchState() {
	fmt.Println("The light switch is:  ", lightSwitchIsOn)
}

func toggleLightSwitch() {
	lightSwitchIsOn = !lightSwitchIsOn
}
