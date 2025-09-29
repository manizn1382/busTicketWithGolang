package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tick/config"
	"tick/model"
)

func AddBus(b model.Bus) (string){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AddBus: ",err)
	}

	defer db.Close()

	res,err := db.Exec(
		`insert into bus 
		(plateNumber,capacity,busType,coId)
		values
		(?,?,?,?)`,
		b.PlateNumber,b.Capacity,b.Type,b.CompanyId,
	)

	if err != nil{
		log.Fatal(err)
		return err.Error()
	}

	id,_ := res.LastInsertId()
	return fmt.Sprintf("%s: %d","last insert id for bus is: ",id) 

}

func GetBusByPlateNumber(pNum string) (*model.Bus,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetBusByPlateNumber : ",err)
	}

	defer db.Close()

	busInfo := model.Bus{}

	res := db.QueryRow(
		"select * from Bus where plateNumber = ?",
		pNum,
	)

	var r int
	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find bus with this plate number")
	}
	
	
	res.Scan(
		&busInfo.BusId,
		&busInfo.PlateNumber,
		&busInfo.Capacity,
		&busInfo.TripId,
		&busInfo.Type,
		&busInfo.CompanyId,
	)

	return &busInfo,nil

}


func UpdateBus(b *model.Bus) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateBus: ",err)
	}

	defer db.Close()


	
	res,err := db.Exec(
	    `update bus 
		set plateNumber = ?, capacity = ?, tripId = ?, busType = ?, coId = ?
		where busId = ?`,
		b.PlateNumber,b.Capacity,b.TripId,b.Type,b.CompanyId,b.BusId,   
	)

	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't update Bus with these info")
	}
	return &res,nil
}


func DeleteBus(pNum int) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteBus: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from bus 
		 where plateNumber = ?`,
		 pNum,   
	)

	if err!=nil{
		return nil,errors.New("can't execute query for plateNumber you give")
	}

	affect,err := res.RowsAffected()

	if affect == 0{
		return nil,errors.New("it doesn't exist bus with this plateNumber")
	}

	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	return &res,err
}


func AllBus() (*[]model.Bus,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AllBus: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from Bus`)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for AllBus func")
	}

	defer res.Close()


	var busList []model.Bus


	for res.Next(){
		var b model.Bus
		if err := res.Scan(&b.BusId,&b.PlateNumber,&b.Capacity,&b.TripId,&b.Type,&b.CompanyId,);err!=nil{
			return nil,err
		}
		busList = append(busList, b)
	}

	return &busList,nil
}




