package service

import (
	"encoding/json"
	"net/http"
	"tick/db"
)

func ViewSeatStatus(r *http.Request, w http.ResponseWriter){

	r.ParseForm()
	seat := r.FormValue("seatNum")

	seatInfo,err := db.GetSeatByNumber(seat)
	
	if err != nil{
		w.Write([]byte ("can't find seat with this number"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	seatInfoJs,err := json.Marshal(seatInfo)

	if err != nil{
		w.Write([]byte ("can't parse seat to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
		w.Write(seatInfoJs)
		w.WriteHeader(http.StatusOK)
}

func ReserveSeat(seatId int) (bool,string,int) {
	// seat := r.URL.Query().Get("seatNum")

	seatInfo,err := db.GetSeatById(seatId)

	if err != nil{
		return false,"can't get seat detail with this seat Id",http.StatusNotFound
	}

	if seatInfo.Status == "reserve"{
		return false,"this seat is already reserved",http.StatusInternalServerError
	}

	seatInfo.Status = "reserve"
	_,err = db.UpdateSeat(seatInfo)

	if err != nil{
		return false,"can't update seat detail with this seat Id",http.StatusInternalServerError
	}
		return true,"reserved seat successfully",http.StatusAccepted
		

}

func MakeFree(seatId int) (bool,string,int){
	//seat := r.URL.Query().Get("seatNum")

	seatInfo,err := db.GetSeatById(seatId)

	if err != nil{
		return false,"can't get seat detail with this seat Id",http.StatusNotFound
	}

	if seatInfo.Status == "free"{
		return false,"this seat is already free",http.StatusInternalServerError
	} 
	seatInfo.Status = "Free"
	_,err = db.UpdateSeat(seatInfo)

	if err != nil{
		return false,"can't update seat detail with this seat id",http.StatusInternalServerError
	}
	return true,"seat is free",http.StatusOK
	
}