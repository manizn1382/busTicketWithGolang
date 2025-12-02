package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
)

func AddBus(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	pNum := r.FormValue("plate")
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

	fmt.Println(res)


	if res != nil{
		w.Write([]byte ("can't add bus with these info"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
		w.Write([]byte ("bus added successfully"))
		w.WriteHeader(http.StatusOK)
}

func RemoveBus(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	pNum := r.FormValue("plate")

	_,err := db.DeleteBus(pNum)

	if err != nil{
		w.Write([]byte ("can't remove bus with these info"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
		w.Write([]byte ("bus removed successfully"))
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


func UpdateBus(r *http.Request, w http.ResponseWriter){

	r.ParseForm()
	pNum := r.FormValue("plate")
	Type := r.FormValue("Type")
	cap,_ := strconv.Atoi(r.FormValue("Capacity"))
	status := r.FormValue("Status")	
	coId,_ := strconv.Atoi(r.FormValue("company"))
	tripId,_ := strconv.Atoi(r.FormValue("trip"))
	busId,_ := strconv.Atoi(r.FormValue("bus"))

	busInfo := model.Bus{
		PlateNumber: pNum,
		BusId: busId,
		CompanyId: coId,
		TripId: tripId,
		Type: Type,
		Capacity: cap,
		Status: status,
	}

	_,err := db.UpdateBus(&busInfo)



	if err != nil {
		w.Write([]byte("can't update bus with these info"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("bus updated successfully."))
	w.WriteHeader(http.StatusOK)
}

func ViewBusInfo(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	plate := r.FormValue("plate")

	coInfo, err := db.GetBusByPlateNumber(plate)

	if err != nil {
		w.Write([]byte("can't find bus with this id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	coInfoJs, err := json.Marshal(coInfo)

	if err != nil {
		w.Write([]byte("can't parse bus to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(coInfoJs)
	w.WriteHeader(http.StatusOK)

}


func BusList(r *http.Request, w http.ResponseWriter) {

	busInfos, err := db.AllBus()

	if err != nil {
		w.Write([]byte("error in fetching Bus."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	busInfosJs, err := json.Marshal(busInfos)

	if err != nil {
		w.Write([]byte("can't convert data to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(busInfosJs)
	w.WriteHeader(http.StatusOK)
}
