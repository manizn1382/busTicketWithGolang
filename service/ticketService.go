package service

import (
	"encoding/json"
	"net/http"
	"tick/db"
	"tick/model"
	"time"
)

func ReserveTicket(r *http.Request, w http.ResponseWriter){

	var ticket model.Ticket

	if err := json.NewDecoder(r.Body).Decode(&ticket);err!=nil{
		w.Write([]byte("invalid json for reserveTicket"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	ticket.BookTime = time.Now()
	ticket.Status = "reserved"

	if err := db.AddTicket(ticket);err!=nil{
		w.Write([]byte("can't add this ticket to db"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	seat,err := db.GetSeatById(ticket.SeatId)

	if err != nil{
		w.Write([]byte("can't find seat with this id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	seat.Status = "reserve"

	_,err = db.UpdateSeat(seat)

	if err != nil{
		w.Write([]byte("can't update seat with this id"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
	w.WriteHeader(http.StatusAccepted)
}

func PrintTicket(r *http.Request, w http.ResponseWriter){
	var ticket model.Ticket

	if err := json.NewDecoder(r.Body).Decode(&ticket);err!=nil{
		w.Write([]byte("invalid json for printTicket"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	ticketInfo,err := db.GetTicketById(ticket.TicketId)

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
	w.WriteHeader(http.StatusAccepted)
}

func CancelTicket(r *http.Request, w http.ResponseWriter){
	var ticket model.Ticket

	if err := json.NewDecoder(r.Body).Decode(&ticket);err!=nil{
		w.Write([]byte("invalid json for cancelTicket"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	ticketInfo,err := db.GetTicketById(ticket.TicketId)

	if err != nil{
		w.Write([]byte("can't find ticket in cancelTicket func."))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ticketInfo.Status = "canceled"
	_,err = db.UpdateTicket(ticketInfo)

	if err != nil{
		w.Write([]byte("can't update ticket in cancelTicket func."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	seatInfo,err := db.GetSeatById(ticketInfo.SeatId)

	if err != nil{
		w.Write([]byte("can't find seat in cancelTicket func."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	seatInfo.Status = "free"

	_,err = db.UpdateSeat(seatInfo)

	if err != nil{
		w.Write([]byte("can't update seat in cancelTicket func."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
	w.WriteHeader(http.StatusAccepted)
}

func ViewUserTicketsHis(r *http.Request, w http.ResponseWriter){
	
}