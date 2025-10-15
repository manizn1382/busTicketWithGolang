package service

import (
	"encoding/json"
	"net/http"
	"tick/db"
	"tick/model"
	"time"
)

func SetPayment(r *http.Request, w http.ResponseWriter){
	
	var p model.Payment
	
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil{
		w.Write([]byte("can't decode request body of payment"))
		w.WriteHeader(http.StatusConflict)
	}else{
		p.CreateAt = time.Now()
		err := db.AddPayment(p)
		if err != nil{
			w.Write([]byte("can't add payment to database"))
			w.WriteHeader(http.StatusInternalServerError)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

//func RedirectToPaymentGateway(){}

func UpdateStatus(r *http.Request, w http.ResponseWriter){
	
	var payment model.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)

	if err != nil{
		w.Write([]byte("can't decode request body of payment"))
		w.WriteHeader(http.StatusConflict)
	}else{
		_,err := db.UpdatePayment(payment)
		if err != nil{
			w.Write([]byte("can't update payment in database"))
			w.WriteHeader(http.StatusInternalServerError)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func RefundPrice(){}