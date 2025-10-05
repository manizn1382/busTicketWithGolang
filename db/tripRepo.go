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

func TripValidation(t model.Trip) (error) {

	statusV,_ := regexp.Compile(`(?i)[closed|open]`)
	originV,_ := regexp.Compile(`^[a-zA-Z]{1,30}`)
	destV,_ := regexp.Compile(`^[a-zA-Z]{1,30}`)

	if !statusV.MatchString(t.Status){return errors.New("status does not match the pattern")}
	if !originV.MatchString(t.Origin){return errors.New("origin does not match the pattern")}
	if !destV.MatchString(t.Dest){return errors.New("dest does not match the pattern")}


	return nil
}


func AddTrip(t model.Trip) (error){



	if err := TripValidation(t);err != nil{
		return err
	}

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		return err
	}

	defer db.Close()

	_,err = db.Exec(
		`insert into trip
		(origin,dest,departureTime,arrivalTime,price,stat,distance)
		values
		(?,?,?,?,?,?,?)`,
		t.Origin,t.Dest,t.DepartureTime,t.ArrivalTime,t.Price,t.Status,t.Distance,
	)

	if err != nil{
		return err
	}

	//id,_ := res.LastInsertId()
	return err 

}



func GetTripById(tId int) (*model.Trip,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetTripById : ",err)
	}

	defer db.Close()

	tripInfo := model.Trip{}

	res := db.QueryRow(
		"select * from trip where tripId = ?",
		tId,
	)
	
	var r int

	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find trip with this trip id")
	}
	
	res.Scan(
		&tripInfo.TripId,
		&tripInfo.Origin,
		&tripInfo.Dest,
		&tripInfo.DepartureTime,
		&tripInfo.ArrivalTime,
		&tripInfo.Price,
		&tripInfo.Status,
		&tripInfo.Distance,
	)

	return &tripInfo,nil

}




func UpdateTrip(t *model.Trip) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateTrip: ",err)
	}

	defer db.Close()


	
	res,err := db.Exec(
	    `update trip 
		set origin = ?, dest = ?, departureTime = ?, arrivalTime = ?, price = ?, stat = ?, distance = ?
		where tripId = ?`,
		t.Origin,t.Dest,t.DepartureTime,t.ArrivalTime,t.Price,t.Status,t.Distance,t.TripId,  
	)

	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't update trip with these info")
	}
	return &res,nil
}


func DeleteTrip(tripId int) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteTrip: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from trip 
		 where tripId = ?`,
		 tripId,   
	)

	if err!=nil{
		return nil,errors.New("can't execute query for tripId you give")
	}

	affect,err := res.RowsAffected()

	if affect == 0{
		return nil,errors.New("it doesn't exist trip with this tripId")
	}

	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	return &res,err
}


func AllTrip() (*[]model.Trip,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AllTrip: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from trip`)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for AllTrip func")
	}

	defer res.Close()


	var tripList []model.Trip


	for res.Next(){
		var t model.Trip
		if err := res.Scan(&t.TripId,&t.Origin,&t.Dest,&t.DepartureTime,&t.ArrivalTime,&t.Price,&t.Status,&t.Distance);err!=nil{
			return nil,err
		}
		tripList = append(tripList, t)
	}

	return &tripList,nil
}




