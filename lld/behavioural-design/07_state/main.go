package main

import "fmt"

type TransportationMode interface {
	calcETA() int
	getDirection() string
}

type Walking struct{}

func (w *Walking) calcETA() int {
	fmt.Println("Calc ETA For Walking")
	return 10
}

func (w *Walking) getDirection() string {
	fmt.Println("Direction for walking")
	return "10"
}

type Cycling struct{}

func (w *Cycling) calcETA() int {
	fmt.Println("Calc ETA For Cycling")
	return 10
}

func (w *Cycling) getDirection() string {
	fmt.Println("Direction for Cycling")
	return "10"
}

type Train struct{}

func (w *Train) calcETA() int {
	fmt.Println("Calc ETA For Train")
	return 10
}

func (w *Train) getDirection() string {
	fmt.Println("Direction for Train")
	return "10"
}

type Car struct{}

func (w *Car) calcETA() int {
	fmt.Println("Calc ETA For Car")
	return 10
}

func (w *Car) getDirection() string {
	fmt.Println("Direction for Car")
	return "10"
}

type DirectionService struct {
	mode TransportationMode
}

func (ds *DirectionService) getETA() int {
	return ds.mode.calcETA()
}

func (ds *DirectionService) getDirection() string {
	return ds.mode.getDirection()
}

func main(){
	directionService := DirectionService{mode: &Car{}}
	directionService.getDirection()
	directionService.getETA()

	directionService = DirectionService{mode: &Train{}}
	directionService.getDirection()
	directionService.getETA()
}