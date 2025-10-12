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
		w.WriteHeader(404)
	}else{
		seatInfo,err := json.Marshal(seatInfo)

		if err != nil{
			w.Write([]byte ("can't parse seat to json"))
			w.WriteHeader(500)
		}else{
			w.Write(seatInfo)
			w.WriteHeader(200)
		}
	}

}

func ReserveSeat(r *http.Request, w http.ResponseWriter){
	seat := r.URL.Query().Get("seatNum")

	seatInfo,err := db.GetSeatByNumber(seat)

	if err != nil{
		w.Write([]byte("can't get seat detail with this seat number"))
		w.WriteHeader(404)
	}else{
		seatInfo.Status = "reserve"
		_,err := db.UpdateSeat(seatInfo)

		if err != nil{
			w.Write([]byte("can't update seat detail with this seat number"))
			w.WriteHeader(500)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(200)
		}
	}

}

func MakeFree(r *http.Request, w http.ResponseWriter){
	seat := r.URL.Query().Get("seatNum")

	seatInfo,err := db.GetSeatByNumber(seat)

	if err != nil{
		w.Write([]byte("can't get seat detail with this seat number"))
		w.WriteHeader(404)
	}else{
		seatInfo.Status = "Free"
		_,err := db.UpdateSeat(seatInfo)

		if err != nil{
			w.Write([]byte("can't update seat detail with this seat number"))
			w.WriteHeader(500)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(200)
		}
	}
}