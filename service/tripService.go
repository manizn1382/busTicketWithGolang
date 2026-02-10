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
		w.Write([]byte("error in parsing data in setTrip in trip service"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	

	tripInfo := model.Trip{
		Origin: origin,
		Dest: dest,
		DepartureTime: departureTime,
		ArrivalTime: arrivalTime,
		Price: float32(price),
		Status: status,
		Distance: float32(dist),
	}


	if err := db.AddTrip(tripInfo);err != nil{
		w.Write([]byte("error in setTrip in trip service"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Trip added successfully"))
	w.WriteHeader(http.StatusOK)
	
}

func SearchByOrigin(r *http.Request,w http.ResponseWriter) {

	r.ParseForm()
	org := r.FormValue("origin")
	tripInfo,err := db.GetTripByOrigin(org)

	if err != nil{
		w.Write([]byte("can't find trip with this origin"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
		tripInfoJs,err := json.Marshal(tripInfo)

		if err != nil{
			w.Write([]byte("error in parsing json in searchByOrigin"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write(tripInfoJs)
			w.WriteHeader(http.StatusOK)
		
}

func SearchByDest(r *http.Request,w http.ResponseWriter) {
	r.ParseForm()
	dest := r.FormValue("destination")
	tripInfo,err := db.GetTripByDest(dest)

	if err != nil{
		w.Write([]byte("can't find trip with this destination in searchBydest"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

		tripInfoJs,err := json.Marshal(tripInfo)

		if err != nil{
			w.Write([]byte("error in parsing json in searchByDest"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write(tripInfoJs)
			w.WriteHeader(http.StatusOK)
		
	
}

func SearchByDate(r *http.Request,w http.ResponseWriter) {
	r.ParseForm()

	date := r.FormValue("date")
	tripInfo,err := db.GetTripByDate(date)


	if err != nil{
		w.Write([]byte("can't find trip with this date in searchBydate"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
		tripInfoJs,err := json.Marshal(tripInfo)

		if err != nil{
			w.Write([]byte("error in parsing json in searchByDate"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write(tripInfoJs)
			w.WriteHeader(http.StatusOK)
		
	
}


func ChangeStatus(r *http.Request,w http.ResponseWriter) {
	 r.ParseForm()

	 tripId,errId := strconv.Atoi(r.FormValue("Id"))
	 status := r.FormValue("stat")

	 if errId != nil{
		w.Write([]byte("can't convert Id to int in changeStatus"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	 }
		 tripInfo,err := db.GetTripById(tripId)

		 if err != nil{
			w.Write([]byte("can't find trip with this id in changeStatus"))
			w.WriteHeader(http.StatusNotFound)
			return
		 }
			tripInfo.Status = status
			_,err = db.UpdateTrip(tripInfo)
			if err != nil{
				w.Write([]byte("can't update trip with this id in change status"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
				w.Write([]byte("success"))
				w.WriteHeader(http.StatusOK)
			
}


func TripUpdate(r *http.Request,w http.ResponseWriter){
	r.ParseForm()

	origin := r.FormValue("origin")
	dest := r.FormValue("dest")
	tripId,err := strconv.Atoi(r.FormValue("tripId"))
	departureTime,err1 := time.Parse(config.RefTime,r.FormValue("departureTime"))
	arrivalTime,err2 := time.Parse(config.RefTime,r.FormValue("arrivalTime"))
	price,err3 := strconv.ParseFloat(r.FormValue("price"),32)
	status := r.FormValue("status")
	dist,err4 := strconv.ParseFloat(r.FormValue("distance"),32)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err != nil{
		w.Write([]byte("error in parsing data in TripUpdate in trip service"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	

	tripInfo := model.Trip{
		TripId: tripId,
		Origin: origin,
		Dest: dest,
		DepartureTime: departureTime,
		ArrivalTime: arrivalTime,
		Price: float32(price),
		Status: status,
		Distance: float32(dist),
	}

	_,err = db.UpdateTrip(&tripInfo)



	if err != nil {
		w.Write([]byte("trip can't update with these info."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	
	w.Write([]byte("trip updated successfully."))
	w.WriteHeader(http.StatusOK)
}


func ViewTripInfo(r *http.Request,w http.ResponseWriter) {
	r.ParseForm()

	tripId := r.FormValue("tripId")
	Id,errId := strconv.Atoi(tripId)

	if errId != nil{
		w.Write([]byte("can't parse id in ViewTripInfo "))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

		trip,err := db.GetTripById(Id)
	
		if err != nil{
			w.Write([]byte("can't find trip with this id in ViewTripInfo"))
			w.WriteHeader(http.StatusNotFound)
			return
		}
			res,err := json.Marshal(trip)
	
			if err != nil{
				w.Write([]byte("can't parse data to json in ViewTripInfo"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
				w.Write(res)
				w.WriteHeader(http.StatusOK)
}