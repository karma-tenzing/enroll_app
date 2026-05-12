package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/date"
	"myapp/utils/httpResp"
	"net/http"
	"strings"
)

// handler function to handle add enrollment
func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll

	decorder := json.NewDecoder(r.Body)
	if err := decorder.Decode(&e); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// fmt.Println(e)
	current_date := date.GetDate()
	e.Date_Enrolled = current_date

	// pass e to model

	saveErr := e.EnrollStud()
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.ResponseWithError(w, http.StatusForbidden, saveErr.Error())
		} else {
			httpResp.ResponseWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	} else {
		httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"status": "Student enrolled"})
	}
	// fmt.Println(saveErr)
}
