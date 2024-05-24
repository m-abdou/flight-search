package handler

import (
	"daemon/handler/Model"
	"fmt"
	"log"
	"os"
	"sync"
)

func Manipulate() map[string][]Model.Flight {
	csvLocation := "./datafiles/"
	entries, err := os.ReadDir(csvLocation)
	if err != nil {
		log.Fatal("cant read data source directory", err)
	}

	var totalResult []Model.Flight

	var resultCan = make(chan []Model.Flight)

	var dataFilter = map[string][]Model.Flight{}

	var wgMain sync.WaitGroup

	for _, e := range entries {
		wgMain.Add(1)

		go func(location string, name string, result chan<- []Model.Flight) {
			defer wgMain.Done()
			//data := ProcessFile(location, name)
			//result <- nil
		}(csvLocation, e.Name(), resultCan)
		//go process(&wg, csvLocation, e.Name(), resultCan)
	}

	wgMain.Wait()

	for flightData := range resultCan {
		totalResult = append(totalResult, flightData...)
	}

	//go func(result <-chan []Model.Flight) {
	//	for flightData := range resultCan {
	//		totalResult = append(totalResult, flightData...)
	//	}
	//
	//	close(resultCan)
	//}(resultCan)

	dataFilter["data"] = totalResult
	for _, flight := range totalResult {
		dataFilter[flight.Airline] = append(dataFilter[flight.Airline], flight)
		dataFilter[flight.DepartureAirport] = append(dataFilter[flight.DepartureAirport], flight)
		dataFilter[flight.DepartureDate] = append(dataFilter[flight.DepartureDate], flight)
		dataFilter[flight.ArrivalAirport] = append(dataFilter[flight.ArrivalAirport], flight)
		dataFilter[flight.ReturnDate] = append(dataFilter[flight.ReturnDate], flight)
	}

	fmt.Println("total result: ", len(totalResult))

	return dataFilter
}

func process(wg *sync.WaitGroup, location string, name string, result chan<- []Model.Flight) {
	defer wg.Done()
	data := ProcessFile(location, name)
	result <- data
}
