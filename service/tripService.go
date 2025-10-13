package service

import (
	"net/http"
	"tick/db"
	"tick/model"
	"time"
	"tick/config"
	"strconv"
)

func SetTrip(r *http.Request,w http.ResponseWriter) {
	r.ParseForm()
	origin := r.FormValue("origin")
	dest := r.FormValue("dest")
	departureTime,err1 := time.Parse(config.RefTime,r.FormValue("departureTime"))
	arrivalTime,err2 := time.Parse(config.RefTime,r.FormValue("arrivalTime"))
	price,err3 := strconv.ParseFloat(r.FormValue("price"),32)
	status := r.FormValue("status")
	dist,err4 := strconv.ParseFloat(r.FormValue("distance"),32)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil{
		w.Write([]byte("error in converting/parsing data in trip service"))
		w.WriteHeader(500)
	}else{
		tripInfo := model.Trip{
			Origin: origin,
			Dest: dest,
			DepartureTime: departureTime,
			ArrivalTime: arrivalTime,
			Price: float32(price),
			Status: status,
			Distance: float32(dist),
		}
		err := db.AddTrip(tripInfo)
		if err != nil{
			w.Write([]byte("error in add trip in trip service"))
			w.WriteHeader(500)
		}else{
			w.Write([]byte("success"))
			w.WriteHeader(200)
		}

	}

}

func SearchByOrigin() {}

func SearchByDest() {}

func SearchByDate() {}

func ChangeStatus() {}

func ViewTripInfo() {}