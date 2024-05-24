package handler

import (
	"context"
	"daemon/handler/Model"
	"fmt"
	"github.com/actforgood/bigcsvreader"
	"strconv"
	"sync"
)

func ProcessFile(csvLocation string, fileName string) []Model.Flight {
	fmt.Println("file name:", fileName)

	bigCSV := bigcsvreader.New()
	bigCSV.SetFilePath(csvLocation + fileName)
	bigCSV.MaxGoroutinesNo = 10
	bigCSV.BufferSize = 81920
	bigCSV.ColumnsDelimiter = '^'

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	var wg sync.WaitGroup

	rowsChans, errsChan := bigCSV.Read(ctx)
	if errsChan != nil {
		fmt.Sprintf("cant read data from file %v ", fileName)
	}

	var result []Model.Flight

	for i := 0; i < len(rowsChans); i++ {
		wg.Add(1)
		go rowWorker(rowsChans[i], &wg, &result)
	}

	//wg.Add(1)
	//go errWorker(errsChan, &wg)

	wg.Wait()

	fmt.Println("result count of file name: ", fileName, len(result))

	return result
}

func processRow(row []string, result *[]Model.Flight) {
	flight := Model.Flight{
		DepartureAirport: row[4],
		DepartureDate:    row[6],
		ArrivalAirport:   row[5],
		ReturnDate:       row[7],
		Airline:          row[len(row)-1],
	}
	price, _ := strconv.ParseFloat(row[12], 64)

	flight.TotalPrice = price

	*result = append(*result, flight)
}
