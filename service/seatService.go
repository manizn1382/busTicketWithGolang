package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
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

func SeatUpdate(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	seatId, seatIdErr := strconv.Atoi(r.FormValue("Id"))
	busId,busIdErr := strconv.Atoi(r.FormValue("busId"))
	seatNum,seatNumErr := strconv.Atoi(r.FormValue("seat"))
	status := r.FormValue("status")
	detail := r.FormValue("detail")

	if seatIdErr != nil {
		w.Write([]byte("seat Id has invalid format."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if busIdErr != nil {
		w.Write([]byte("bus Id has invalid format."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if seatNumErr != nil {
		w.Write([]byte("company Id has invalid format."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	seatInfo := model.Seat{
		SeatId: seatId,
		BusId: busId,
		SeatNum: seatNum,
		Status: status,
		Description: detail,
	}

	_, err := db.UpdateSeat(&seatInfo)

	if err != nil {
		w.Write([]byte("can't update seat in system."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("seat updated successfully."))
	w.WriteHeader(http.StatusOK)
}


func CreateSeat(r *http.Request, w http.ResponseWriter){

	r.ParseForm()
	busId,busIdErr := strconv.Atoi(r.FormValue("busId"))
	seatNum,seatNumErr := strconv.Atoi(r.FormValue("seat"))
	status := r.FormValue("status")
	detail := r.FormValue("detail")

	if busIdErr != nil {
		w.Write([]byte("bus Id has invalid format."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if seatNumErr != nil {
		w.Write([]byte("company Id has invalid format."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	seatInfo := model.Seat{
		BusId: busId,
		SeatNum: seatNum,
		Status: status,
		Description: detail,
	}

	err := db.AddSeat(seatInfo)


	if err != nil {
		w.Write([]byte("can't Add seat in system."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("seat Created successfully."))
	w.WriteHeader(http.StatusOK)
}

func SeatList(r *http.Request, w http.ResponseWriter) {

	seatInfos, err := db.AllSeat()

	if err != nil {
		w.Write([]byte("error in fetching seats."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	seatInfosJs, err := json.Marshal(seatInfos)

	if err != nil {
		w.Write([]byte("can't convert data to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(seatInfosJs)
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