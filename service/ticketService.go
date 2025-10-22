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

func PrintTicket(){}

func CancelTicket(){}

func ViewUserTicketsHis(){}