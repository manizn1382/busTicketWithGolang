package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
	"time"
)

func ReserveTicket(r *http.Request, w http.ResponseWriter){

	r.ParseForm()

	tId,err1 := strconv.Atoi(r.FormValue("tripId"))
	uId,err2 := strconv.Atoi(r.FormValue("userId"))
	sId,err3 := strconv.Atoi(r.FormValue("seatId"))

	if err1 != nil || err2 != nil || err3 != nil{
		w.Write([]byte("error in converting numeric Id from string to int"))
		w.WriteHeader(http.StatusConflict)
		return
	}


	success,reason,code := ReserveSeat(sId)

	if (!success){
		w.Write([]byte(reason))
		w.WriteHeader(code)
		return
	}

	ticketInfo := model.Ticket{
		TripId: tId,
		UserId: uId,
		SeatId: sId,
		Status: "reserved",
		BookTime: time.Now(),
	}

	tripInfo,err := db.GetTripById(tId)

	if err != nil {
		w.Write([]byte("can't find trip with this id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if tripInfo.Status == "closed"{
		w.Write([]byte("this trip is closed. select another trip."))
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err := db.AddTicket(ticketInfo);err!=nil{
		fmt.Println(err)
		w.Write([]byte("can't add this ticket to db"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	w.Write([]byte("ticket added successfully"))
	w.WriteHeader(http.StatusOK)
}

func PrintTicket(r *http.Request, w http.ResponseWriter){

	r.ParseForm()

	tId,err := strconv.Atoi(r.FormValue("ticketId"))


	if err != nil{
		w.Write([]byte("error in converting ticket Id to numeric value"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	ticketInfo,err := db.GetTicketById(tId)

	if err != nil{
		w.Write([]byte("can't find ticket in PrintTicket func."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ticketJson,err := json.Marshal(ticketInfo)

	if err != nil{
		w.Write([]byte("can't convert ticket to json in printTicket."))
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Write(ticketJson)
	w.WriteHeader(http.StatusOK)
}

func CancelTicket(r *http.Request, w http.ResponseWriter){

	r.ParseForm()

	tId,err := strconv.Atoi(r.FormValue("ticketId"))

	if err != nil{
		w.Write([]byte("error in converting ticketId to int value"))
		w.WriteHeader(http.StatusConflict)
		return
	}


	ticketInfo,err := db.GetTicketById(tId)

	if err != nil{
		w.Write([]byte("can't find ticket in cancelTicket func."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	success,reason,code := RefundPrice(ticketInfo) 
	
	if(!success){
		w.Write([]byte(reason))
		w.WriteHeader(code)
		return
	}
	
	success,reason,code = MakeFree(ticketInfo.SeatId)
	
	if(!success){
		w.Write([]byte(reason))
		w.WriteHeader(code)
		return
	}
	

	ticketInfo.Status = "canceled"
	_,err = db.UpdateTicket(ticketInfo)

	if err != nil{
		w.Write([]byte("can't update ticket in cancelTicket func."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ticket cancelled successfully"))
	w.WriteHeader(http.StatusOK)
}

func ViewUserTicketsHis(r *http.Request, w http.ResponseWriter){

	r.ParseForm()

	uId,err := strconv.Atoi(r.FormValue("userId"))

	if err != nil {
		w.Write([]byte("can't convert the string value to int value"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	ticketList,err := db.GetUserTicketHis(uId)

	if err != nil{
		w.Write([]byte("can't find ticket in viewUserTicketHis func."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ticketListJson,err := json.Marshal(ticketList)

	if err != nil{
		w.Write([]byte("can't parse ticketList into json in viewUserTicketHis func."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(ticketListJson)
	w.WriteHeader(http.StatusOK)

}