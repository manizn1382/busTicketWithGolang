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
		w.Write([]byte (res.Error()))
	}else{
		w.Write([]byte ("ok"))
		w.WriteHeader(200)
	}


}

func ViewCompanyInfo(r *http.Request, w http.ResponseWriter){

	r.ParseForm()
	phone := r.FormValue("phone")

	coInfo,err := db.GetCompanyByPhone(phone)
	
	if err != nil{
		w.Write([]byte ("can't find company with this phone"))
		w.WriteHeader(404)
	}else{
		coInfo,err := json.Marshal(coInfo)

		if err != nil{
			w.Write([]byte ("can't parse company to json"))
			w.WriteHeader(500)
		}else{
			w.Write(coInfo)
			w.WriteHeader(200)
		}
	}
}

func CompanyList(r *http.Request, w http.ResponseWriter){

	coInfos,err := db.AllCo()

	if err != nil{
		w.Write([]byte("error in fetching companies."))
		w.WriteHeader(404)
	}else{

		coInfos,err := json.Marshal(coInfos)

		if err != nil{
			w.Write([]byte("can't convert data to json"))
			w.WriteHeader(500)
		}else{
			w.Write(coInfos)
			w.WriteHeader(200)
		}
		
	}

}