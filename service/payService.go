package service

import (
	"encoding/json"
	"net/http"
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
			w.WriteHeader(http.StatusOK)
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

func RefundPrice(ticket *model.Ticket) (bool,string,int) {


	payInfo, err := db.GetPayByTicketId(ticket.TicketId)

	if err != nil {
		return false,"can't find payment with this ticket id",http.StatusNotFound
	}

	
	if payInfo.PayStatus != "paid" {
		return true,"",http.StatusAccepted
	}



	tripInfo, err := db.GetTripById(ticket.TripId)

	if err != nil {
		return false,"can't find trip with this trip id",http.StatusNotFound
	}


	refPercent := calculateRefundPercent(tripInfo)

	refAmount := payInfo.Amount * refPercent

	PayToUser(ticket.UserId, refAmount)


	payInfo.PayStatus = "Canceled"
	

	_,err = db.UpdatePayment(*payInfo)

	if err != nil{
		return false,"can't update pay with this info in refund price",http.StatusInternalServerError
	}

	return true,"",http.StatusAccepted
	
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
