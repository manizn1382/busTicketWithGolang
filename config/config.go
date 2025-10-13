package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


var DBUser string 
var DBPass string
var DBHost string
var DBPort string
var DBName string
var RefTime string
var Dsn    string

func SetEnv() {

    err := godotenv.Load(".env")
    
    if err != nil {
        log.Fatal("Error loading .env file")
    }

        DBUser = os.Getenv("userName")
        DBPass = os.Getenv("passWord")
        DBHost = os.Getenv("host")
        DBPort = os.Getenv("port")
        DBName = os.Getenv("dbName")
        Dsn = os.Getenv("dsn")
        RefTime = os.Getenv("refTime")

}
