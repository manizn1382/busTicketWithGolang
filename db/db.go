package db

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"log"
	"tick/config"
)

func AddUser(userId int, userName string, phoneNumber string) (string) {
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db: ",err)
	}
	defer db.Close()

	res,err := db.Exec(
		"insert into User (userId,userName,phoneNumber) values (?, ?, ?)",
		userId,userName,phoneNumber,
	)

	if err != nil{
		log.Fatal(err)
	}
	id,_ := res.LastInsertId()
	return fmt.Sprintf("%s: %d","last insert id is: ",id) 
	

}

func GetUserById() {

}