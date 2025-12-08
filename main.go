package main

import (
	"tick/config"
	"tick/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	
	config.SetEnv()
	
	r := gin.Default()

    bus := r.Group("/bus")
    {
    bus.Any("/:operation", controller.BusHandler)
    }

    company := r.Group("/company")
    {
    company.Any("/:operation", controller.CoHandler)
    }

    seat := r.Group("/seat")
    {
    seat.Any("/:operation", controller.SeatHandler)
    }


    r.Run(":8081")


}
