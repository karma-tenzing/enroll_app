package model

import (
	"myapp/dataStore/postgres"
)

type Enroll struct {
	StdID         int64  `json:stdid`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "INSERT INTO enroll (std_id, course_id, date_enrolled) VALUES ($1, $2, $3);"

func (e *Enroll) EnrollStud() error {
	// fmt.Println("modal in e", e)

	_, err := postgres.Db.Exec(queryEnrollStd, e.StdID, e.CourseID, e.Date_Enrolled)
	return err

}
