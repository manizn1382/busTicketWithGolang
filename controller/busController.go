package controller

import (
	"net/http"
	"tick/service"

	"github.com/gin-gonic/gin"
)

func BusHandler(c *gin.Context) {

	op := c.Param("operation")

	switch op {
	case "CreateBus":
		service.AddBus(c.Request, c.Writer)
	case "BindBusToTrip":
		service.BindBusToTrip(c.Request, c.Writer)
	case "BindBusToCompany":
		service.BindBusToCompany(c.Request, c.Writer)
	case "DeleteBus":
		service.RemoveBus(c.Request, c.Writer)
	case "BusList":
		service.BusList(c.Request, c.Writer)
	case "UpdateBus":
		service.UpdateBus(c.Request, c.Writer)
	case "ViewBusInfo":
		service.ViewBusInfo(c.Request, c.Writer)
	default:
		{
			c.Writer.Write([]byte(c.Request.URL.Path))
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
