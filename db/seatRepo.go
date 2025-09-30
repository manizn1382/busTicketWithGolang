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



func AddSeat(s model.Seat) (string){


	if err := SeatValidation(s);err!=nil{
		return err.Error()
	} 

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AddSeat: ",err)
	}

	defer db.Close()

	res,err := db.Exec(
		`insert into seat 
		(busId,seatNum,stat,detail)
		values
		(?,?,?,?)`,
		s.BusId,s.SeatNum,s.Status,s.Description,
	)

	if err != nil{
		log.Fatal(err)
		return err.Error()
	}

	id,_ := res.LastInsertId()
	return fmt.Sprintf("%s: %d","last insert id for bus is: ",id) 

}

func GetSeatById(Id string) (*model.Seat,error){
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




