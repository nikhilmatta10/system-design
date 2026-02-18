package main

import "fmt"

type TransportationMode string

const (
	WALKING TransportationMode = "walking"
	CYCLING TransportationMode = "cycling"
	CAR     TransportationMode = "car"
	TRAIN   TransportationMode = "train"
)

type DirectionService struct {
	mode TransportationMode
}

func NewDirectionService(mode TransportationMode) *DirectionService {
	return &DirectionService{mode: mode}
}

func (ds *DirectionService) setMode(mode TransportationMode) {
	ds.mode = mode
}

func (ds *DirectionService) getETA() int {
	switch ds.mode {
	case WALKING:
		fmt.Println("Calculating ETA for walking 10")
		return 10
	case CYCLING:
		fmt.Println("Calc ETA for cycling 5")
		return 5
	case CAR:
		fmt.Println("Calc ETA for car 2")
		return 2
	case TRAIN:
		fmt.Println("Calc ETA for train 3")
		return 3
	default:
		fmt.Println("Unkown Mode")
	}
	return 0
}

func (ds *DirectionService) getDirection() string {
	switch ds.mode {
	case WALKING:
		fmt.Println("Calculating ETA for walking 10")
		return "10"
	case CYCLING:
		fmt.Println("Calc ETA for cycling 5")
		return "05"
	case CAR:
		fmt.Println("Calc ETA for car 2")
		return "02"
	case TRAIN:
		fmt.Println("Calc ETA for train 3")
		return "03"
	default:
		fmt.Println("Unkown Mode")
	}
	return "0"
}

func main() {
	directionService := NewDirectionService(TRAIN)
	directionService.setMode(CAR)

	directionService.getDirection()
	directionService.getETA()
}

