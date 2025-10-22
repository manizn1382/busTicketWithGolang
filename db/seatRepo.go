package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tick/config"
	"tick/model"
	"regexp"
)


func SeatValidation(s model.Seat) (error) {

	statusV,_ := regexp.Compile(`(?i)[reserve|free]`)

	if !statusV.MatchString(s.Status){return errors.New("status does not match the pattern")}


	return nil
}



func AddSeat(s model.Seat) (error){


	if err := SeatValidation(s);err!=nil{
		return err
	} 

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		return err
	}

	defer db.Close()

	_,err = db.Exec(
		`insert into seat 
		(busId,seatNum,stat,detail)
		values
		(?,?,?,?)`,
		s.BusId,s.SeatNum,s.Status,s.Description,
	)

	if err != nil{
		return err
	}

	//id,_ := res.LastInsertId()
	return nil 

}

func GetSeatById(Id int) (*model.Seat,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetSeatById : ",err)
	}

	defer db.Close()

	SeatInfo := model.Seat{}

	res := db.QueryRow(
		"select * from seats where seatId = ?",
		Id,
	)

	var r int
	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find seat with this id")
	}
	
	
	res.Scan(
		&SeatInfo.SeatId,
		&SeatInfo.BusId,
		&SeatInfo.SeatNum,
		&SeatInfo.Status,
		&SeatInfo.Description,
	)

	return &SeatInfo,nil

}

func GetSeatByNumber(sNum string) (*model.Seat,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetSeatByNumber : ",err)
	}

	defer db.Close()

	SeatInfo := model.Seat{}

	res := db.QueryRow(
		"select * from seats where seatNum = ?",
		sNum,
	)

	var r int
	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find seat with this seat Number")
	}
	
	
	res.Scan(
		&SeatInfo.SeatId,
		&SeatInfo.BusId,
		&SeatInfo.SeatNum,
		&SeatInfo.Status,
		&SeatInfo.Description,
	)

	return &SeatInfo,nil

}


func UpdateSeat(s *model.Seat) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateSeat: ",err)
	}

	defer db.Close()


	
	res,err := db.Exec(
	    `update seats 
		set busId = ?, seatNum = ?, stat = ?, detail = ?
		where seatId = ?`,
		s.BusId,s.SeatNum,s.Status,s.Description,s.SeatId,   
	)

	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't update Bus with these info")
	}
	return &res,nil
}


func AllSeat() (*[]model.Seat,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AllBus: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from seats`)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for AllSeat func")
	}

	defer res.Close()


	var seatList []model.Seat


	for res.Next(){
		var s model.Seat
		if err := res.Scan(&s.SeatId,&s.BusId,&s.SeatNum,&s.Status,&s.Description,);err!=nil{
			return nil,err
		}
		seatList = append(seatList, s)
	}

	return &seatList,nil
}




