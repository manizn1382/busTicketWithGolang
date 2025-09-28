package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tick/config"
	"tick/model"
)

func AddTrip(t model.Trip) (string){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AddTrip: ",err)
	}

	defer db.Close()

	res,err := db.Exec(
		`insert into trip
		(origin,dest,departureTime,arrivalTime,price,stat,distance)
		values
		(?,?,?,?,?,?,?)`,
		t.Origin,t.Dest,t.DepartureTime,t.ArrivalTime,t.Price,t.Status,t.Distance,
	)

	if err != nil{
		log.Fatal(err)
		return err.Error()
	}

	id,_ := res.LastInsertId()
	return fmt.Sprintf("%s: %d","last insert id for trip is: ",id) 

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
	).Scan(
		&tripInfo.TripId,
		&tripInfo.Origin,
		&tripInfo.Dest,
		&tripInfo.DepartureTime,
		&tripInfo.ArrivalTime,
		&tripInfo.Price,
		&tripInfo.Status,
		&tripInfo.Distance,
	)


	if res == nil{
		return nil,errors.New("can't execute query with given tripId")
	}

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




