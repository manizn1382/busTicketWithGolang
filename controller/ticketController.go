package controller

import (
	"tick/service"

	"github.com/gin-gonic/gin"
)

func TicketHandler(c *gin.Context){

	op := c.Param("operation")

	switch op{
	case "ReserveTicket":
		service.ReserveTicket(c.Request,c.Writer)
	case "PrintTicket":
		service.PrintTicket(c.Request,c.Writer)
	case "CancelTicket":
		service.CancelTicket(c.Request,c.Writer)
	case "ViewUserTicketsHis":
		service.ViewUserTicketsHis(c.Request,c.Writer)			
	}
}