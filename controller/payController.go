package controller

import (
	"tick/service"

	"github.com/gin-gonic/gin"
)

func PayHandler(c *gin.Context){
	op := c.Param("operation")

	switch op{
	case "SetPayment":
		service.SetPayment(c.Request,c.Writer)
	case "UpdateStatus":
		service.UpdateStatus(c.Request,c.Writer)	
	}
}