package service

import (
	"net/http"
	"strconv"
	"tick/model"
	"tick/db"
)

func AddBus(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	pNum := r.FormValue("plateNumber")
	Type := r.FormValue("Type")
	cap,_ := strconv.Atoi(r.FormValue("Capacity"))


	bus := model.Bus{
		PlateNumber: pNum,
		Capacity: cap,
		Type: Type,
	}

	res := db.AddBus(bus)

	if res != nil{
		w.Write([]byte (res.Error()))
	}else{
		w.Write([]byte ("ok"))
		w.WriteHeader(200)
	}

}

func BindBusToTrip(r *http.Request, w http.ResponseWriter) {
	
	r.ParseForm()
	plateNumber := r.FormValue("plate")
	tripdId,_ := strconv.Atoi(r.FormValue("trip"))

	busInfo,err := db.GetBusByPlateNumber(plateNumber)

	if err != nil{
		w.Write([]byte("it doesn't exist bus with this plate number"))
		w.WriteHeader(404)
	}else{
		busInfo.TripId = tripdId
		_,err := db.UpdateBus(busInfo)
		if err != nil{
			w.Write([]byte("can't update with this tripId"))
			w.WriteHeader(500)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(200)
		}
	}

}

func BindBusToCompany(r *http.Request, w http.ResponseWriter) {
	
	r.ParseForm()
	plateNumber := r.FormValue("plate")
	coId,_ := strconv.Atoi(r.FormValue("company"))

	busInfo,err := db.GetBusByPlateNumber(plateNumber)

	if err != nil{
		w.Write([]byte("it doesn't exist bus with this plate number"))
		w.WriteHeader(404)
	}else{
		busInfo.CompanyId = coId
		_,err := db.UpdateBus(busInfo)
		if err != nil{
			w.Write([]byte("can't update with this coId"))
			w.WriteHeader(500)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(200)
		}
	}
}
