package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tick/config"
	"tick/db"
	"tick/model"
	"time"
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

func SearchByOrigin(r *http.Request,w http.ResponseWriter) {
	org := r.URL.Query().Get("origin")
	tripInfo,err := db.GetTripByOrigin(org)

	if err != nil{
		w.Write([]byte("trip doesn't exist"))
		w.WriteHeader(404)
	}else{
		tripInfo,err := json.Marshal(tripInfo)

		if err != nil{
			w.Write([]byte("error in parsing json"))
			w.WriteHeader(500)
		}else{
			w.Write(tripInfo)
			w.WriteHeader(200)
		}
	}
}

func SearchByDest(r *http.Request,w http.ResponseWriter) {
	dest := r.URL.Query().Get("destination")
	tripInfo,err := db.GetTripByDest(dest)

	if err != nil{
		w.Write([]byte("trip doesn't exist"))
		w.WriteHeader(404)
	}else{
		tripInfo,err := json.Marshal(tripInfo)

		if err != nil{
			w.Write([]byte("error in parsing json"))
			w.WriteHeader(500)
		}else{
			w.Write(tripInfo)
			w.WriteHeader(200)
		}
	}
}

func SearchByDate(r *http.Request,w http.ResponseWriter) {
	date := r.URL.Query().Get("date")
	tripInfo,err := db.GetTripByDate(date)

	if err != nil{
		w.Write([]byte("trip doesn't exist"))
		w.WriteHeader(404)
	}else{
		tripInfo,err := json.Marshal(tripInfo)

		if err != nil{
			w.Write([]byte("error in parsing json"))
			w.WriteHeader(500)
		}else{
			w.Write(tripInfo)
			w.WriteHeader(200)
		}
	}
}

func ChangeStatus(r *http.Request,w http.ResponseWriter) {
	 r.ParseForm()
	 tripId,errId := strconv.Atoi(r.FormValue("Id"))
	 if errId != nil{
		w.Write([]byte("can't convert Id to int"))
		w.WriteHeader(500)
	 }else{
		 status := r.FormValue("stat")
		 tripInfo,err := db.GetTripById(tripId)
		 if err != nil{
			w.Write([]byte("can't find trip with this id"))
			w.WriteHeader(404)
		 }else{
			tripInfo.Status = status
			_,err := db.UpdateTrip(tripInfo)
			if err != nil{
				w.Write([]byte("can't update trip with this id"))
				w.WriteHeader(500)
			}else{
				w.Write([]byte("success"))
				w.WriteHeader(200)
			}

		 }

	 }

}

func ViewTripInfo(r *http.Request,w http.ResponseWriter) {
	tripId := r.URL.Query().Get("tripId")
	Id,errId := strconv.Atoi(tripId)
	if errId != nil{
		w.Write([]byte("can't parse id"))
		w.WriteHeader(500)
	}else{
		trip,err := db.GetTripById(Id)
	
		if err != nil{
			w.Write([]byte("trip Not Found"))
			w.WriteHeader(404)
		}else{
			res,err := json.Marshal(trip)
	
			if err != nil{
				w.Write([]byte("can't parse data to json"))
				w.WriteHeader(500)
			}else{
				w.Write(res)
				w.WriteHeader(200)
			}
		}
	}

}