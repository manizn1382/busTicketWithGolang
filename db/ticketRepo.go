package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tick/config"
	"tick/model"
	"time"
	"regexp"
)

func TicketValidation(t model.Ticket) (error) {

	statusV,_ := regexp.Compile(`(?i)[reserved|available|canceled]`)

	if !statusV.MatchString(t.Status){return errors.New("status does not match the pattern")}


	return nil
}



func AddTicket(t model.Ticket) (error){

	if err := TicketValidation(t);err!=nil{
		return err
	}


	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		return err
	}

	defer db.Close()

	_,err = db.Exec(
		`insert into ticket
		(tripId,userId,seatId,bookTime,stat)
		values
		(?,?,?,?,?)`,
		t.TripId,t.UserId,t.SeatId,time.Now(),t.Status,
	)

	if err != nil{
		return err
	}

	//id,_ := res.LastInsertId()
	return nil 

}



func GetTicketByTripId(tId int) (*model.Ticket,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetTicketByTripId : ",err)
	}

	defer db.Close()

	ticketInfo := model.Ticket{}

	res := db.QueryRow(
		"select * from ticket where tripId = ?",
		tId,
	)
	
	var r int 
	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find ticket with this trip id")
	}
	
	
	res.Scan(
		&ticketInfo.TicketId,
		&ticketInfo.TripId,
		&ticketInfo.UserId,
		&ticketInfo.SeatId,
		&ticketInfo.BookTime,
		&ticketInfo.Status,
	)

	return &ticketInfo,nil

}


func GetTicketById(tId int) (*model.Ticket,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetTicketById : ",err)
	}

	defer db.Close()

	ticketInfo := model.Ticket{}

	res := db.QueryRow(
		"select * from ticket where ticketId = ?",
		tId,
	)
	
	var r int 
	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find ticket with this id")
	}
	
	
	res.Scan(
		&ticketInfo.TicketId,
		&ticketInfo.TripId,
		&ticketInfo.UserId,
		&ticketInfo.SeatId,
		&ticketInfo.BookTime,
		&ticketInfo.Status,
	)

	return &ticketInfo,nil

}


func UpdateTicket(t *model.Ticket) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateTicket: ",err)
	}

	defer db.Close()


	
	res,err := db.Exec(
	    `update ticket 
		set tripId = ?, userId = ?, seatId = ?, stat = ?
		where ticketId = ?`,
		t.TripId,t.UserId,t.SeatId,t.Status,t.TicketId,  
	)

	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't update ticket with these info")
	}
	return &res,nil
}


func DeleteTicket(ticketId int) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteTicket: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from ticket 
		 where ticketId = ?`,
		 ticketId,   
	)

	if err!=nil{
		return nil,errors.New("can't execute query for ticketId you give")
	}

	affect,err := res.RowsAffected()

	if affect == 0{
		return nil,errors.New("it doesn't exist ticket with this ticketId")
	}

	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	return &res,err
}


func AllTicket() (*[]model.Ticket,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AllTicket: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from ticket`)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for AllTicket func")
	}

	defer res.Close()


	var ticketList []model.Ticket


	for res.Next(){
		var t model.Ticket
		if err := res.Scan(&t.TicketId,&t.TripId,&t.UserId,&t.SeatId,&t.BookTime,&t.Status,);err!=nil{
			return nil,err
		}
		ticketList = append(ticketList, t)
	}

	return &ticketList,nil
}

func GetUserTicketHis(uid int) (*[]model.Ticket,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db GetUserTicketHis: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from ticket where userId = ?`,uid,)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for GetUserTicketHis func")
	}

	defer res.Close()


	var ticketList []model.Ticket


	for res.Next(){
		var t model.Ticket
		if err := res.Scan(&t.TicketId,&t.TripId,&t.UserId,&t.SeatId,&t.BookTime,&t.Status,);err!=nil{
			return nil,err
		}
		ticketList = append(ticketList, t)
	}

	return &ticketList,nil
}




