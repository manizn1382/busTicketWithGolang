package service

import (
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
	"time"
)

func SetPayment(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()

	tId, err := strconv.Atoi(r.FormValue("ticketId"))
	amount, err2 := strconv.Atoi(r.FormValue("amount"))
	payType := r.FormValue("payType")
	payStatus := r.FormValue("payStatus")

	if err == nil || err2 == nil {
		w.Write([]byte("can't convert form values to numeric value"))
		w.WriteHeader(http.StatusConflict)
		return
	}

	payInfo := model.Payment{
		TicketId:  tId,
		PayType:   payType,
		Amount:    float32(amount),
		PayStatus: payStatus,
	}

	if err = db.AddPayment(payInfo); err != nil {
		w.Write([]byte("can't add payment"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payInfo.CreateAt = time.Now()

	w.Write([]byte("payment added successfully"))
	w.WriteHeader(http.StatusOK)

}

//func RedirectToPaymentGateway(){}

func UpdateStatus(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()

	payId,err := strconv.Atoi(r.FormValue("payId"))


	if err != nil{
		w.Write([]byte("can't convert payId to numeric value in UpdateStatus"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payInfo,err := db.GetPayById(payId)

	if err != nil{
		w.Write([]byte("can't find payment with these info"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payInfo.PayStatus = "Cancelled"

	
	_, err = db.UpdatePayment(*payInfo)

	if err != nil {
		w.Write([]byte("can't update payment with these info"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("payment updated successfully"))
	w.WriteHeader(http.StatusAccepted)
	
}

func RefundPrice(ticket *model.Ticket) (bool, string, int) {

	payInfo, err := db.GetPayByTicketId(ticket.TicketId)

	if err != nil {
		return true, "can't find payment with this ticket id", http.StatusNotFound
	}

	if payInfo.PayStatus != "paid" {
		return true, "", http.StatusAccepted
	}

	tripInfo, err := db.GetTripById(ticket.TripId)

	if err != nil {
		return true, "can't find trip with this trip id", http.StatusNotFound
	}

	refPercent := calculateRefundPercent(tripInfo)

	refAmount := payInfo.Amount * refPercent

	PayToUser(ticket.UserId, refAmount)

	payInfo.PayStatus = "Canceled"

	_, err = db.UpdatePayment(*payInfo)

	if err != nil {
		return false, "can't update pay with this info in refund price", http.StatusInternalServerError
	}

	return true, "", http.StatusAccepted

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
