package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tick/config"
	"tick/model"
	"time"
	"regexp"
)

func PayValidation(p model.Payment) (error) {

	typeV,_ := regexp.Compile(`(?i)[card|cash]`)
	statusV,_ := regexp.Compile(`(?i)[complete|inProgress|failed]`)

	if !typeV.MatchString(p.PayType){return errors.New("pay type does not match the pattern")}
	if !statusV.MatchString(p.PayStatus){return errors.New("status does not match the pattern")}


	return nil
}


func AddPayment(p model.Payment) (string){

	if err:= PayValidation(p);err!=nil{
		return err.Error()
	}


	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AddPayment: ",err)
	}

	defer db.Close()

	res,err := db.Exec(
		`insert into  payment
		(ticketId,amount,payType,payStatus,createdAt)
		values
		(?,?,?,?,?)`,
		p.TicketId,p.Amount,p.PayType,p.PayStatus,time.Now(),
	)

	if err != nil{
		log.Fatal(err)
		return err.Error()
	}

	id,_ := res.LastInsertId()
	return fmt.Sprintf("%s: %d","last insert id for payment is: ",id) 

}



func GetPayByTicketId(tId int) (*model.Payment,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db in GetPayByTicketId : ",err)
	}

	defer db.Close()

	payInfo := model.Payment{}

	res := db.QueryRow(
		"select * from payment where ticketId = ?",
		tId,
	)

	var r int
	rowErr := res.Scan(&r)


	if rowErr == sql.ErrNoRows{
		return nil,errors.New("can't find payment by this ticket id")
	}
	
	
	res.Scan(
		&payInfo.PaymentId,
		&payInfo.TicketId,
		&payInfo.Amount,
		&payInfo.PayType,
		&payInfo.PayStatus,
		&payInfo.CreateAt,
	)

	return &payInfo,nil

}




func UpdatePayment(p *model.Payment) (*sql.Result,error) {

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db UpdatePayment: ",err)
	}

	defer db.Close()


	
	res,err := db.Exec(
	    `update payment 
		set ticketId = ?, amount = ?, payType = ?, payStatus = ?
		where payId = ?`,
		p.TicketId,p.Amount,p.PayType,p.PayStatus,p.PaymentId,  
	)

	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't update Payment with these info")
	}
	return &res,nil
}


func DeletePayment(pId int) (*sql.Result,error){
	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db DeletePayment: ",err)
	}

	defer db.Close()

	
	res,err := db.Exec(
	    `delete from payment 
		 where payId = ?`,
		 pId,   
	)

	if err!=nil{
		return nil,errors.New("can't execute query for payId you give")
	}

	affect,err := res.RowsAffected()

	if affect == 0{
		return nil,errors.New("it doesn't exist payment with this payId")
	}

	if err != nil{
		log.Fatal(err)
		return nil,err
	}

	return &res,err
}


func AllPayment() (*[]model.Payment,error){

	db,err := sql.Open("mysql",config.Dsn)

	if err != nil{
		fmt.Println("error opening db AllPayment: ",err)
	}

	defer db.Close()

	res,err := db.Query(`select * from payment`)

	
	if err != nil{
		log.Fatal(err)
		return nil,errors.New("can't execute query for AllPayment func")
	}

	defer res.Close()


	var payList []model.Payment


	for res.Next(){
		var p model.Payment
		if err := res.Scan(&p.PaymentId,&p.TicketId,&p.Amount,&p.PayType,&p.PayStatus,&p.CreateAt,);err!=nil{
			return nil,err
		}
		payList = append(payList, p)
	}

	return &payList,nil
}




