package service
import (
	"net/http"
	"tick/model"
	"tick/db"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"encoding/json"
)


func SignIn(r *http.Request, w http.ResponseWriter) {
	
	r.ParseForm()
	phone := r.FormValue("phone")
	passWord := r.FormValue("passWord")

	user,err := db.GetUserByPhone(phone)

	if err != nil{
		w.Write([]byte(err.Error()))
		w.WriteHeader(204)
	}

	hash := sha256.Sum256([]byte(passWord))
	PassWordHash := hex.EncodeToString(hash[:])
	
	if PassWordHash != user.PassHash{
		w.Write([]byte("password is incorrect"))
		w.WriteHeader(204)
	}else{
		w.Write([]byte("ok"))
		w.WriteHeader(200)	
	}

}



func SignUp(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	userName := r.FormValue("userName")
	phone := r.FormValue("phone")
	nationalId := r.FormValue("nationalId")
	passWord := r.FormValue("passWord")

	userInfo := model.User {
		Name: userName,
		Phone: phone,
		NationalId: nationalId,
		PassHash: passWord,
	}

	resp := db.AddUser(userInfo)

	if resp != nil{
		w.Write([]byte("can't signUp"))
		w.WriteHeader(204)
	}

	w.Write([]byte("ok"))
	w.WriteHeader(200)

}



func EditProfile(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()

	userId,err := strconv.Atoi(r.FormValue("Id"))

	if err != nil{
		w.Write([]byte("can't convert Id to int"))
		w.WriteHeader(204)
	}else{

		userInfo := model.User{
		Name: r.FormValue("userName"),
		Role: r.FormValue("Role"),
		PassHash: r.FormValue("PassWord"),
		Phone: r.FormValue("phone"),
		NationalId: r.FormValue("nationalId"),
		UserId: userId,
		}

		_,err := db.UpdateUser(&userInfo)

		if err != nil{
			w.Write([]byte("can't update user"))
			w.WriteHeader(204)
		}else{
			w.Write([]byte("ok"))
			w.WriteHeader(200)
		}
	}
	
}


func ViewProfile(r *http.Request, w http.ResponseWriter) {

	phone := r.URL.Query().Get("phone")

	user,err := db.GetUserByPhone(phone)

	if err != nil{
		w.Write([]byte("user Not Found"))
		w.WriteHeader(404)
	}else{
		res,err := json.Marshal(user)

		if err != nil{
			w.Write([]byte("can't parse data to json"))
			w.WriteHeader(500)
		}else{
			w.Write(res)
			w.WriteHeader(200)
		}
	}

}




