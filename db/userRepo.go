package db

import(
	"tick/config"
	"fmt"
	"time"
	"log"
	"tick/model"
	"database/sql"
)

func AddUser(userName , phoneNumber string) (string) {
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println(err)
	}

	defer db.Close()

	res,err := db.Exec(
		"insert into User (userName,phoneNumber,createTime) values (?, ?, ?)",
	    userName,phoneNumber,time.Now(),
	)

	if err != nil{
		log.Fatal(err)
	}

	id,_ := res.LastInsertId()
	return fmt.Sprintf("%s: %d","last insert id is: ",id) 

}


func GetUserByPhone(phoneNumber string) (*model.User,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetUserByPhone : ",err)
	}

	defer db.Close()

	userInfo := model.User{}

	res := db.QueryRow(
		"select * from User where phoneNumber = ?",
		phoneNumber,
	).Scan(
		&userInfo.UserId,
		&userInfo.Name,
		&userInfo.Role,
		&userInfo.PassHash,
		&userInfo.Phone,
		&userInfo.CreateAt,
		&userInfo.NationalId,
	)


	if res != nil{
		return nil,res
	}

	return &userInfo,nil

}

func GetUserById(userId int) (*model.User,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db: ",err)
	}

	defer db.Close()

	userInfo := model.User{}

	res := db.QueryRow(
		"select * from User where userId = ?",
		userId,
	).Scan(
		&userInfo.UserId,
		&userInfo.CreateAt,
		&userInfo.Name,
		&userInfo.NationalId,
		&userInfo.Phone,
		&userInfo.PassHash,
		&userInfo.Role,
	)

	fmt.Println(res)


	if res != nil{
		return nil,res
	}

	
	return &userInfo,nil

}

func UpdateUser(userInfo *model.User) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateUser: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `update User 
		set userName = ?, userRole = ?, phoneNumber = ?, passwordHash = ?, nationalId = ?
		where userId = ?`,
		userInfo.Name,userInfo.Role,userInfo.Phone,12,userInfo.NationalId,userInfo.UserId,   
	)

	if err != nil{
		log.Fatal(err)
		return nil,err
	}
	return &res,nil
}


func DeleteUser(userId int) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteUser: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from User 
		 where userId = ?`,
		 userId,   
	)

	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	return &res,err
}



func AllUser() (*[]model.User,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteUser: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from User`)

	
	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	defer res.Close()

	var users []model.User

	for res.Next(){
		var u model.User
		if err := res.Scan(&u.UserId,&u.CreateAt,&u.Name,&u.NationalId,&u.Phone,&u.PassHash,&u.Role);err!=nil{
			return nil,err
		}
		users = append(users, u)
	}

	return &users,nil
}