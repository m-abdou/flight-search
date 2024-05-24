package handler

import (
	"daemon/handler/Model"
	"github.com/actforgood/bigcsvreader"
	"sync"
)

func rowWorker(rowsChan bigcsvreader.RowsChan, waitGr *sync.WaitGroup, resultChan *[]Model.Flight) {
	defer waitGr.Done()

	for row := range rowsChan {
		processRow(row, resultChan)
	}
}

func errWorker(errsChan bigcsvreader.ErrsChan, waitGr *sync.WaitGroup) {
	defer waitGr.Done()

	for err := range errsChan {
		handleError(err)
	}
}
