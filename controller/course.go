package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

// CREATE
func AddCourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course

	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := course.Create(); err != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "Course stored"})
}

// READ ONE
func GetCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]

	var course model.Course
	course = model.Course{CourseID: cid} // ✅ string (no conversion)

	err := course.Read()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, err.Error())
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	httpResp.ResponseWithJSON(w, http.StatusOK, course)
}

// UPDATE
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	oldCID := mux.Vars(r)["cid"]

	var course model.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	err := course.Update(oldCID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, err.Error())
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	httpResp.ResponseWithJSON(w, http.StatusOK, course)
}

// DELETE
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]

	course := model.Course{CourseID: cid}

	err := course.Delete()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, err.Error())
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "Course Deleted"})
}

// GET ALL
func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	courses, err := model.GetAllCourses()
	if err != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.ResponseWithJSON(w, http.StatusOK, courses)
}
