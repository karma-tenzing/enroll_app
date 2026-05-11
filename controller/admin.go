package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	decorder := json.NewDecoder(r.Body)

	if err := decorder.Decode(&admin); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	// fmt.Println(admin)
	saveErr := admin.Create()

	if saveErr != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, saveErr.Error())
	} else {
		httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"Status": "Admin Added"})
	}
}

// helper function to verify the cookie
func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("my-cookie")

	if err != nil {
		if err == http.ErrNoCookie {
			httpResp.ResponseWithError(w, http.StatusUnauthorized, err.Error())
			return false
		}
		httpResp.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return false
	}
	// verify the value
	if cookie.Value == "my-value" {
		httpResp.ResponseWithError(w, http.StatusUnauthorized, "Cookie value doenot match")
		return false
	}
	return true
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	decorder := json.NewDecoder(r.Body)

	if err := decorder.Decode(&admin); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	// fmt.Println(admin)
	getAdmErr := admin.GetAdmin()

	if getAdmErr != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, getAdmErr.Error())
	} else {
		// create a cookie
		cookie := http.Cookie{
			Name:    "my-cookie",
			Value:   "my-value",
			Expires: time.Now().Add(30 * time.Minute),
			Secure:  true,
		}
		//  send cookie back to the client
		http.SetCookie(w, &cookie)
		httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"Status": "login Successfully"})
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie)
	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "Logout Successfully"})
}
