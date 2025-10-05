package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tick/config"
	"tick/model"
	"regexp"
)

func companyValidation(c model.Company) (error) {

	nameV,_ := regexp.Compile(`^[a-zA-Z]{1,25}`)
	phoneV,_ := regexp.Compile(`\+98[0-9]{10}`)

	if !nameV.MatchString(c.Name){return errors.New("name does not match the pattern")}
	if !phoneV.MatchString(c.SupportPhone){return errors.New("phone does not match the pattern")}


	return nil
}



func AddCompany(c model.Company) (error){

	if err:= companyValidation(c);err!=nil{
		return err
	}

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		return err
	}

	defer db.Close()

	_,err = db.Exec(
		`insert into company 
		(coName,supportPhone,coAddress)
		values
		(?,?,?)`,
		c.Name,c.SupportPhone,c.Address,
	)

	if err != nil{
		return err
	}

	// id,_ := res.LastInsertId()
	return nil 

}

func GetCompanyByPhone(phone string) (*model.Company,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetCompanyByPhone : ",err)
	}

	defer db.Close()

	coInfo := model.Company{}

	res := db.QueryRow(
		"select * from company where supportPhone = ?",
		phone,
	)

	var r int
	rowErr := res.Scan(&r)

	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find dompany with this phoneNumber")
	}
	
	
	res.Scan(
		&coInfo.CompanyId,
		&coInfo.Name,
		&coInfo.SupportPhone,
		&coInfo.Address,
	)

	return &coInfo,nil

}


func UpdateCo(c *model.Company) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdateCo: ",err)
	}

	defer db.Close()


	
	res,err := db.Exec(
	    `update company 
		set coName = ?, supportPhone = ?, coAddress = ?
		where companyId = ?`,
		c.Name,c.SupportPhone,c.Address,c.CompanyId,   
	)

	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't update company with these info")
	}
	return &res,nil
}


func DeleteCo(coName string) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeleteCo: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from company 
		 where coName = ?`,
		 coName,   
	)

	if err!=nil{
		return nil,errors.New("can't execute query for company name you give")
	}

	affect,err := res.RowsAffected()

	if affect == 0{
		return nil,errors.New("it doesn't exist company with this Name")
	}

	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	return &res,err
}


func AllCo() (*[]model.Company,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AllCo: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from company`)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for AllCo func")
	}

	defer res.Close()


	var companies []model.Company


	for res.Next(){
		var c model.Company
		if err := res.Scan(&c.CompanyId,&c.Name,&c.SupportPhone,&c.Address,);err!=nil{
			return nil,err
		}
		companies = append(companies, c)
	}

	return &companies,nil
}




