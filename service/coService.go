package service

import (
	"net/http"
	_"strconv"
	"tick/model"
	"tick/db"
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

func ViewCompanyInfo(){}

func CompanyList(){}