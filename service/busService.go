package service

import (
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
)

func AddBus(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	pNum := r.FormValue("plateNumber")
	Type := r.FormValue("Type")
	cap,_ := strconv.Atoi(r.FormValue("Capacity"))
	status := r.FormValue("Status")


	bus := model.Bus{
		PlateNumber: pNum,
		Capacity: cap,
		Type: Type,
		Status: status,
	}

	res := db.AddBus(bus)


	if res != nil{
		w.Write([]byte ("can't add bus with these info"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
		w.Write([]byte ("bus added successfully"))
		w.WriteHeader(http.StatusOK)
}

func BindBusToTrip(r *http.Request, w http.ResponseWriter) {
	
	r.ParseForm()
	plateNumber := r.FormValue("plate")
	tripdId,_ := strconv.Atoi(r.FormValue("trip"))

	busInfo,err := db.GetBusByPlateNumber(plateNumber)

	if err != nil{
		w.Write([]byte("it doesn't exist bus with this plate number"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
		busInfo.TripId = tripdId
		_,err = db.UpdateBus(busInfo)
		if err != nil{
			w.Write([]byte("can't update with this tripId"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write([]byte("bus binded to trip successfully"))
			w.WriteHeader(http.StatusOK)
}

func BindBusToCompany(r *http.Request, w http.ResponseWriter) {
	
	r.ParseForm()
	plateNumber := r.FormValue("plate")
	coId,_ := strconv.Atoi(r.FormValue("company"))

	busInfo,err := db.GetBusByPlateNumber(plateNumber)

	if err != nil{
		w.Write([]byte("it doesn't exist bus with this plate number"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
		busInfo.CompanyId = coId
		_,err = db.UpdateBus(busInfo)
		
		if err != nil{
			w.Write([]byte("can't update with this coId"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write([]byte("bus binded to company successfully"))
			w.WriteHeader(http.StatusOK)
		
	
}
