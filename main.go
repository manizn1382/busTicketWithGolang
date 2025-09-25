package main

import (
	"fmt"
	"tick/config"
	"tick/db"
	_ "github.com/go-sql-driver/mysql"
	//_"reflect"
)






func main() {
	
	config.SetEnv()
	
	message,err := db.AllUser()

	fmt.Println(*message,err)

	// val := reflect.ValueOf(message)
    // typ := reflect.TypeOf(message)

    // for i := 0; i < val.NumField(); i++ {
    //     field := typ.Field(i)
    //     value := val.Field(i)
    //     fmt.Printf("%s: %v\n", field.Name, value)
    // }


}
