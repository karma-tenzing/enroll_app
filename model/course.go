package model

import "myapp/dataStore/postgres"

type Course struct {
	CourseID   string `json:"cid"`
	CourseName string `json:"cname"`
}

// SQL Queries
const queryInsertCourse = "INSERT INTO course(cid, coursename) VALUES ($1, $2);"
const queryGetCourse = "SELECT cid, coursename FROM course WHERE cid = $1;"
const queryUpdateCourse = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid;"
const queryDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid;"
const queryGetAllCourse = "SELECT cid, coursename FROM course;"

// CREATE
func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertCourse, c.CourseID, c.CourseName)
	return err
}

// READ
func (c *Course) Read() error {
	row := postgres.Db.QueryRow(queryGetCourse, c.CourseID)
	return row.Scan(&c.CourseID, &c.CourseName)
}

// UPDATE
func (c *Course) Update(oldCID string) error {
	row := postgres.Db.QueryRow(queryUpdateCourse, c.CourseID, c.CourseName, oldCID)
	return row.Scan(&c.CourseID)
}

// DELETE
func (c *Course) Delete() error {
	return postgres.Db.QueryRow(queryDeleteCourse, c.CourseID).Scan(&c.CourseID)
}

// GET ALL
func GetAllCourses() ([]Course, error) {
	rows, err := postgres.Db.Query(queryGetAllCourse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.CourseID, &c.CourseName); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}
