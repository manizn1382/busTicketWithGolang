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
}

func BindBusToCompany(r *http.Request, w http.ResponseWriter) {}
