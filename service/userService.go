package service

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"tick/db"
	"tick/model"
)


func SignIn(r *http.Request, w http.ResponseWriter) {
	
	r.ParseForm()
	phone := r.FormValue("phone")
	passWord := r.FormValue("passWord")

	user,err := db.GetUserByPhone(phone)


	if err != nil{
		w.Write([]byte("can't find user with this info"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	hash := sha256.Sum256([]byte(passWord))
	PassWordHash := hex.EncodeToString(hash[:])

	
	if PassWordHash != user.PassHash{
		w.Write([]byte("password is incorrect"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Signed In successfully"))
	w.WriteHeader(http.StatusOK)	
	
}



func SignUp(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()
	userName := r.FormValue("userName")
	phone := r.FormValue("phone")
	nationalId := r.FormValue("nationalId")
	passWord := r.FormValue("passWord")
	Role := r.FormValue("Role")

	userInfo := model.User {
		Name: userName,
		Phone: phone,
		NationalId: nationalId,
		PassHash: passWord,
		Role: Role,
	}

	resp := db.AddUser(userInfo)

	if resp != nil{
		w.Write([]byte("signUp failed with these info"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("user Signed Up successfully"))
	w.WriteHeader(http.StatusOK)

}



func EditProfile(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()

	userId,err := strconv.Atoi(r.FormValue("Id"))

	if err != nil{
		w.Write([]byte("can't convert Id to int in editProfile"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



		userInfo := model.User{
		Name: r.FormValue("userName"),
		Role: r.FormValue("Role"),
		PassHash: r.FormValue("passWord"),
		Phone: r.FormValue("phone"),
		NationalId: r.FormValue("nationalId"),
		UserId: userId,
		}

		_,err = db.UpdateUser(&userInfo)

		if err != nil{
			w.Write([]byte("can't update user in editProfile"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write([]byte("user Updated successfully"))
			w.WriteHeader(http.StatusOK)
	
}


func ViewProfile(r *http.Request, w http.ResponseWriter) {

	r.ParseForm()

	phone := r.FormValue("phone")

	user,err := db.GetUserByPhone(phone)


	if err != nil{
		w.Write([]byte("user Not Found with this phone number"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
		res,err := json.Marshal(user)

		if err != nil{
			w.Write([]byte("can't parse data to json in viewProfile"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
			w.Write(res)
			w.WriteHeader(http.StatusOK)
}

func DeleteUser(r *http.Request,w http.ResponseWriter){
	r.ParseForm()

	nationalId := r.FormValue("nationalId")

	_,err := db.DeleteUser(nationalId)


	if err != nil{
		w.Write([]byte("can't delete User with this national Id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
			w.Write([]byte("User removed successfully"))
			w.WriteHeader(http.StatusOK)
}

func UserList(r *http.Request, w http.ResponseWriter) {

	userInfos, err := db.AllUser()

	if err != nil {
		w.Write([]byte("error in fetching users."))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userInfosJs, err := json.Marshal(userInfos)

	if err != nil {
		w.Write([]byte("can't convert data to json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(userInfosJs)
	w.WriteHeader(http.StatusOK)
}



