package router

import (
	"daemon/controller/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (r *Router) search(c *gin.Context) {
	departure := c.GetString("departure")
	destination := c.GetString("destination")
	departureDate := c.GetString("departure_date")
	arrivalDate := c.GetString("arrival_date")
	page := c.GetString("page")
	perPage := c.GetString("per_page")
	sortDirection := c.GetString("sort_direction")
	airlines := c.GetString("airlines[]")

	response := r.controller.SearchDetails(request.Request{
		Departure:     departure,
		DepartureDate: departureDate,
		Destination:   destination,
		ArrivalDate:   arrivalDate,
		Page:          page,
		PerPage:       perPage,
		SortDirection: sortDirection,
		Airlines:      strings.Split(airlines, ","),
	})

	c.JSON(http.StatusOK, response)
}
