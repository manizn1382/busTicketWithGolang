package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)






func main() {
	
	err := godotenv.Load()

	if err != nil{
		fmt.Println("error occured.")
	}

	username  := os.Getenv("userName")
	password  := os.Getenv("passWord")
	host      := os.Getenv("host")
	port      := os.Getenv("port")
	dbName        := os.Getenv("dbName")
	dsn       := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	db,err := sql.Open("mysql",dsn)

	if err != nil{
		fmt.Print(err)
	}

	defer db.Close()


}
