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

    

    r.Run(":8081")

// 	bus := r.Group("/bus")
// {
//     bus.POST("/create", controller.CreateBus)
//     bus.POST("/update-status", controller.UpdateBusStatus)
//     bus.POST("/assign-driver", controller.AssignDriver)
//     bus.POST("/add-seats", controller.AddSeats)
//     bus.POST("/remove-seats", controller.RemoveSeats)
//     bus.POST("/search", controller.SearchBus)
// }

	// val := reflect.ValueOf(message)
    // typ := reflect.TypeOf(message)

    // for i := 0; i < val.NumField(); i++ {
    //     field := typ.Field(i)
    //     value := val.Field(i)
    //     fmt.Printf("%s: %v\n", field.Name, value)
    // }


}
