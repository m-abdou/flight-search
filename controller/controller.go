package controller

import (
	"daemon/controller/request"
	"daemon/handler/Model"
	"daemon/service"
)

type Controller struct {
	data map[string][]Model.Flight
}

func Init(data map[string][]Model.Flight) *Controller {
	return &Controller{
		data: data,
	}
}

func (c *Controller) SearchDetails(request request.Request) service.SearchFilter {
	return service.Search(request, c.data)
}
