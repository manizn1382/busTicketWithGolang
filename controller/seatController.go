package controller

import (
	"net/http"
	"tick/service"

	"github.com/gin-gonic/gin"
)

func SeatHandler(c *gin.Context) {

	op := c.Param("operation")

	switch op {
	case "ViewSeatStatus":
		service.ViewSeatStatus(c.Request, c.Writer)
	case "CreateSeat":
		service.CreateSeat(c.Request, c.Writer)
	case "SeatUpdate":
		service.SeatUpdate(c.Request,c.Writer)	
	case "SeatList":
		service.SeatList(c.Request,c.Writer)		
	default:
		{
			c.Writer.Write([]byte(c.Request.URL.Path))
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
