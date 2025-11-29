package controller

import (
	"fmt"
	"net/http"
	"tick/service"

	"github.com/gin-gonic/gin"
)



func BusHandler(c *gin.Context) {

	op := c.Param("operation")

	fmt.Println(op)

	switch(op){
	case "CreateBus":service.AddBus(c.Request,c.Writer)
	case "BindBusToTrip":service.BindBusToTrip(c.Request,c.Writer)
	case "BindBusToCompany":service.BindBusToCompany(c.Request,c.Writer)
	default: {
		c.Writer.Write([]byte(c.Request.URL.Path))
		c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}








