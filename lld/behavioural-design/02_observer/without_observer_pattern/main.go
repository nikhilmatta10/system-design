package main

import "fmt"

type DisplayDevice struct{}

func (dd *DisplayDevice) showTemperature(temp float64) {
	fmt.Printf("Temp on this device is %f \n", temp)
}

type WeatherStation struct {
	temperature   float64
	displayDevice DisplayDevice // can be multiple such devices later on
}

func NewWeatherStation(displayDevice DisplayDevice) (*WeatherStation) {
	return &WeatherStation{displayDevice: displayDevice}
}

func (ws *WeatherStation) setTemperature(temp float64) {
	ws.temperature = temp
	ws.notifyDevice()
}

func (ws *WeatherStation) notifyDevice() {
	ws.displayDevice.showTemperature(ws.temperature)
}


func main() {
	displayDevice := DisplayDevice{}
	weatherStation := NewWeatherStation(displayDevice)

	weatherStation.setTemperature(26)
	weatherStation.setTemperature(32)
}