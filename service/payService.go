package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
	"time"
)

func SetPayment(r *http.Request, w http.ResponseWriter) {

	var p model.Payment

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		w.Write([]byte("can't decode request body of payment"))
		w.WriteHeader(http.StatusConflict)
	} else {
		p.CreateAt = time.Now()
		err := db.AddPayment(p)
		if err != nil {
			w.Write([]byte("can't add payment to database"))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("success"))
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

//func RedirectToPaymentGateway(){}

func UpdateStatus(r *http.Request, w http.ResponseWriter) {

	var payment model.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)

	if err != nil {
		w.Write([]byte("can't decode request body of payment"))
		w.WriteHeader(http.StatusConflict)
	} else {
		_, err := db.UpdatePayment(payment)
		if err != nil {
			w.Write([]byte("can't update payment in database"))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("success"))
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func RefundPrice(r *http.Request, w http.ResponseWriter) {
	tId, err := strconv.Atoi(r.URL.Query().Get("ticketId"))
	if err != nil {
		w.Write([]byte("can't parse data to int"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	payInfo, err := db.GetPayByTicketId(tId)

	if payInfo.PayStatus != "paid" {
		return
	}

	if err != nil {
		w.Write([]byte("can't find payment with this ticket id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ticketInfo, err := db.GetTicketById(tId)

	if err != nil {
		w.Write([]byte("can't find ticket with this ticket id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tripInfo, err := db.GetTripById(ticketInfo.TripId)

	if err != nil {
		w.Write([]byte("can't find trip with this trip id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	seatInfo,err := db.GetSeatById(ticketInfo.SeatId)

	if err != nil{
		w.Write([]byte("can't find seat with this seat id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}


	refPercent := calculateRefundPercent(tripInfo)

	refAmount := payInfo.Amount * refPercent

	PayToUser(ticketInfo.UserId, refAmount)


	payInfo.PayStatus = "Canceled"
	ticketInfo.Status = "canceled"
	seatInfo.Status = "free"
	

	_,err = db.UpdatePayment(*payInfo)

	if err != nil{
		w.Write([]byte("can't update pay with this info in refund price"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	_,err = db.UpdateTicket(ticketInfo)

	if err != nil{
		w.Write([]byte("can't update ticket with this info in refund price"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_,err = db.UpdateSeat(seatInfo)

	if err != nil{
		w.Write([]byte("can't update seat with this info in refund price"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("refund process paid successfully"))
	w.WriteHeader(http.StatusAccepted)
	
}

func PayToUser(uId int, amount float32) {}

func calculateRefundPercent(trip *model.Trip) float32 {
	diffTime := time.Since(trip.DepartureTime)

	if diffTime.Minutes() <= 25 {
		return 0.9
	}
	if diffTime.Minutes() >= 26 && diffTime.Minutes() <= 50 {
		return 0.5
	} else if diffTime.Hours() < 2 {
		return 0.4
	} else {
		return 0
	}
}
