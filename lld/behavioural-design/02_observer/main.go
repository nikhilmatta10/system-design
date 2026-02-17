package main

import "fmt"

type Observer interface {
	update(temp float64)
}

type Subject interface {
	attachObserver(o Observer)
	detachObserver(o Observer)
	notifyObservers()
}

type WeatherStation struct {
	observerList []Observer
	temperature  float64
}

func (ws *WeatherStation) attachObserver(o Observer) {
	ws.observerList = append(ws.observerList, o)
}

func (ws *WeatherStation) detachObserver(o Observer) {
	for i, observer := range ws.observerList {
		if observer == o {
			ws.observerList = append(ws.observerList[:i], ws.observerList[i+1:]...)
			break
		}
	}
}

func (ws *WeatherStation) notifyObservers() {
	for _, observer := range ws.observerList {
		observer.update(ws.temperature)
	}
}

type DisplayDevice struct{}

func (dd *DisplayDevice) update(temp float64) {
	fmt.Printf("Temperature on Display device is %f \n", temp)
}

type MobileDevice struct{}

func (dd *MobileDevice) update(temp float64) {
	fmt.Printf("Temperature on Mobile Device is %f \n", temp)
}

func main() {
	ws := WeatherStation{}

	displayDevice := DisplayDevice{}
	mobileDevice := MobileDevice{}

	ws.attachObserver(&displayDevice)
	ws.attachObserver(&mobileDevice)

	ws.notifyObservers()

	ws.detachObserver(&mobileDevice)

	ws.notifyObservers()
}