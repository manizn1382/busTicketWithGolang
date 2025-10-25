package service

import (
	"encoding/json"
	"net/http"
	_ "strconv"
	"tick/db"
	"tick/model"
)

func AddCompany(r *http.Request, w http.ResponseWriter){
	
	r.ParseForm()
	supportPhone := r.FormValue("phone")
	coName := r.FormValue("name")
	coAddress := r.FormValue("Addr")

	coInfo := model.Company{
		SupportPhone: supportPhone,
		Name: coName,
		Address: coAddress,
	}

	res := db.AddCompany(coInfo)

	if res != nil{
		w.Write([]byte ("company can't add to system."))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
		w.Write([]byte ("company added successfully."))
		w.WriteHeader(http.StatusOK)

}

func ViewCompanyInfo(r *http.Request, w http.ResponseWriter){

	r.ParseForm()
	phone := r.FormValue("phone")

	coInfo,err := db.GetCompanyByPhone(phone)
	
	if err != nil{
		w.Write([]byte ("can't find company with this phone"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
		coInfoJs,err := json.Marshal(coInfo)

		if err != nil{
			w.Write([]byte ("can't parse company to json"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write(coInfoJs)
			w.WriteHeader(http.StatusOK)
		
}

func CompanyList(r *http.Request, w http.ResponseWriter){

	coInfos,err := db.AllCo()

	if err != nil{
		w.Write([]byte("error in fetching companies."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

		coInfosJs,err := json.Marshal(coInfos)

		if err != nil{
			w.Write([]byte("can't convert data to json"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write(coInfosJs)
			w.WriteHeader(http.StatusOK)
}