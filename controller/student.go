package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	// create the variable of type student to store student info
	var stud model.Student

	// extract the data from the request body sent by client
	jsonObj := json.NewDecoder(r.Body)

	// store jsn data in stud varibale, converting json data to Go object
	err := jsonObj.Decode(&stud)

	if err != nil {
		// sending the response error back to client
		w.Write([]byte(err.Error()))
	}
	r.Body.Close()
	// no error
	// call model and pass studnet info
	dbErr := stud.Create()

	if dbErr != nil {
		w.Write([]byte(dbErr.Error()))
	}
	// no error
	w.Write([]byte("Successfully stored"))

}

// convert string sid to int
func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func GetUserID(userID string) (int64, error) {
	intID, err := strconv.ParseInt(userID, 10, 64)
	return intID, err
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}

	myMap := mux.Vars(r)
	stdid := myMap["sid"]
	stdID, idErr := GetUserID((stdid))

	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	// no error
	var studDet model.Student
	studDet = model.Student{StdId: stdID}
	//pass student data to model
	getErr := studDet.Read()

	//check error
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, getErr.Error())
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, studDet)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	old_Sid := mux.Vars(r)["sid"]
	old_StdId, idErr := getUserId(old_Sid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	updateErr := stud.Update(old_StdId)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, updateErr.Error())
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.ResponseWithJSON(w, http.StatusOK, stud)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	sid := mux.Vars(r)["sid"]
	stdID, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// way 1 for method
	stud := model.Student{StdId: stdID}
	delErr := stud.Delete()
	if delErr != nil {
		switch delErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, delErr.Error())
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, delErr.Error())
		}
	} else {
		httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "Student Deleted"})
	}

	// way 2
	// delErr := Delete(stdID)
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	students, getErr := model.GetAllStudents()
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
	} else {
		httpResp.ResponseWithJSON(w, http.StatusOK, students)
	}

}
