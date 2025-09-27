package db

import(
	"tick/config"
	"fmt"
	"time"
	"log"
	"tick/model"
	"database/sql"
	"crypto/sha256"
    "encoding/hex"
)

func AddUser(u model.User) (string) {
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println(err)
	}

	defer db.Close()

	hash := sha256.Sum256([]byte(u.PassHash))
	newPassWordHash := hex.EncodeToString(hash[:])
	u.PassHash = newPassWordHash

	res,err := db.Exec(
		`insert into User 
		(userName,phoneNumber,createTime,passwordHash,nationalId,userRole)
		 values 
		(?, ?, ?, ?, ?, ?)`,
	    u.Name,u.Phone,time.Now(),u.PassHash,u.NationalId,u.Role,
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

func GetUserByNationalId(nId string) (*model.User,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db: ",err)
	}

	defer db.Close()

	userInfo := model.User{}

	res := db.QueryRow(
		"select * from User where nationalId = ?",
		nId,
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

func UpdateUser(u *model.User) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateUser: ",err)
	}

	defer db.Close()

	hash := sha256.Sum256([]byte(u.PassHash))
	newPassWordHash := hex.EncodeToString(hash[:])
	u.PassHash = newPassWordHash


	
	res,err := db.Exec(
	    `update User 
		set userName = ?, userRole = ?, phoneNumber = ?, passwordHash = ?, nationalId = ?
		where userId = ?`,
		u.Name,u.Role,u.Phone,u.PassHash,u.NationalId,u.UserId,   
	)

	if err != nil{
		log.Fatal(err)
		return nil,err
	}
	return &res,nil
}


func DeleteUser(nId int) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteUser: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from User 
		 where nationalId = ?`,
		 nId,   
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
		if err := res.Scan(&u.UserId,&u.Name,&u.Role,&u.PassHash,&u.Phone,&u.CreateAt,&u.NationalId,);err!=nil{
			return nil,err
		}
		users = append(users, u)
	}

	return &users,nil
}