package service

import (
	"daemon/controller/request"
	"daemon/handler/Model"
	"sort"
)

type SearchFilter struct {
	TotalCount  int
	TotalPages  int
	FlightMatch []Model.Flight
	Airline     map[string]Airline
}

type Airline struct {
	Count       int
	LowerPrice  float64
	HigherPrice float64
}

func Search(request request.Request, dataStored map[string][]Model.Flight) SearchFilter {
	var searchFilter SearchFilter

	searchFilter.TotalCount = len(dataStored["data"])
	searchFilter.TotalPages = len(dataStored["data"]) / 50

	if len(request.Airlines) > 0 {
		airlineData := map[string]Airline{}
		for _, airline := range request.Airlines {
			if val, ok := dataStored[airline]; ok {
				low, high := fetchLowAndHighPrice(val)

				airlineData[airline] = Airline{
					Count:       len(val),
					LowerPrice:  low,
					HigherPrice: high,
				}
			}
		}

		searchFilter.Airline = airlineData
	}

	return searchFilter
}

func fetchLowAndHighPrice(flights []Model.Flight) (float64, float64) {
	var price []float64

	for _, flight := range flights {
		price = append(price, flight.TotalPrice)
	}

	sort.Float64s(price)

	return price[0], price[len(price)-1]
}
