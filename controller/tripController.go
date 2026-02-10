package controller

import (
	"tick/service"

	"github.com/gin-gonic/gin"
)

func TripHandler(c *gin.Context){
	op := c.Param("operation")

	switch op{
	case "SetTrip":
		service.SetTrip(c.Request,c.Writer)
	case "SearchByOrigin":
		service.SearchByOrigin(c.Request,c.Writer)
	case "SearchByDest":
		service.SearchByDest(c.Request,c.Writer)
	case "SearchByDate":
		service.SearchByDate(c.Request,c.Writer)
	case "ChangeStatus":
		service.ChangeStatus(c.Request,c.Writer)
	case "ViewTripInfo":
		service.ViewTripInfo(c.Request,c.Writer)
	case "TripUpdate":
		service.TripUpdate(c.Request,c.Writer)					
	}
}