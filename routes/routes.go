package routes

import (
	"fmt"
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {

	// creating a router, mux
	router := mux.NewRouter()

	// register routes in mux router

	//student routes

	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/all", controller.GetAllStudent)
	router.HandleFunc("/student/{sid}", controller.GetStudent).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStudent).Methods("DELETE")

	router.HandleFunc("/course/all", controller.GetAllCourse).Methods("GET")
	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")

	// sign up and login
	router.HandleFunc("/signup", controller.SignUp).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout)

	// enroll api
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")

	// load static files
	fhandler := http.FileServer(http.Dir("./view"))
	//serve static files as a route by registering all stattic files in mux router
	router.PathPrefix("/").Handler(fhandler)

	fmt.Println("Server started successfully")
	// start http server
	log.Fatal(http.ListenAndServe(":8080", router))

	// var port = 8080
	// router := mux.NewRouter()
	// log.Println("Application is running on port", port)
	// log.Fatal(http.ListenAndServe(":8080", router))
}
